from selenium import webdriver
from selenium.webdriver.common.by import By
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC
from datetime import datetime
from pathlib import Path
import os
import json

PATH = os.path.dirname(os.path.abspath(__file__))
URL = 'https://dl.acm.org/loi/cacm/group/'

years_data_urls = []
journals = []

driver = webdriver.Chrome()
driver.get(URL)
driver.implicitly_wait(2)
dec_panel = driver.find_element(By.CLASS_NAME, 'scroll')
dec_elems = dec_panel.find_elements(By.TAG_NAME, 'li')

# берем все пары decade.year
for du in driver.find_elements(By.CLASS_NAME, 'scroll')[1:len(dec_elems)+1]:
    for year in du.find_elements(By.TAG_NAME, 'li'):
        years_data_urls.append(year.find_element(By.TAG_NAME, "a").get_attribute("data-url"))
tab_content_counter = 1

# и идем по каждому году
for data_url in years_data_urls:
    current_year_journals = []
    driver.get(f'{URL}{data_url}')

    # убираем окно с куки
    if len(driver.find_elements(By.ID, 'CybotCookiebotDialogBody')) != 0:
        WebDriverWait(driver, 1).until(EC.element_to_be_clickable((By.ID, 'CybotCookiebotDialogBodyLevelButtonLevelOptinDeclineAll'))).click()
    
    # ищем список журналов текущего года
    while len(driver.find_elements(By.CLASS_NAME, 'tab__content')[tab_content_counter].find_elements(By.TAG_NAME, 'li')) == 0:
        tab_content_counter += 1
    content = driver.find_elements(By.CLASS_NAME, 'tab__content')[tab_content_counter].find_elements(By.TAG_NAME, 'li')
    
    # заполняем предварительный список журналов текущего года
    for journal in content:
        current_year_journals.append({'href': journal.find_element(By.TAG_NAME, 'a').get_attribute('href'),
                                     'title': f'{journal.find_element(By.CLASS_NAME, "coverDate").text}, {journal.find_element(By.CLASS_NAME, "issue").text}',
                                     'decade': data_url[1:5],
                                     'year': data_url[7:11]})
    
    # заходим на каждую страницу журнала, записываем ссылки на все доступные статьи
    for sp_journal in current_year_journals:
        driver.get(sp_journal['href'])
        journals_elems = driver.find_elements(By.CLASS_NAME, 'issue-item-container')
        articles = []
        for j_el in journals_elems:
            file_name = j_el.find_element(By.CLASS_NAME, 'issue-item__content-right').find_element(By.TAG_NAME, 'a').text
            file_details = j_el.find_element(By.CLASS_NAME, 'issue-item__detail').find_element(By.TAG_NAME, 'a').text
            articles.append({file_name: file_details})
        sp_journal['articles'] = articles
    # сохраняем статьи и прикрепляем их к журналу
    journals.append(current_year_journals)
    break

result = json.dumps(journals, indent=4)

# запись в файл
with open(f'{Path(PATH).parent.absolute()}/puller/files/{datetime.now().strftime("%H%M%S%f")}.json', 'w') as file:
    file.write(result)

driver.close()