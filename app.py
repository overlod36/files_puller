from selenium import webdriver
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.common.by import By
from selenium.webdriver.support import expected_conditions as EC
from selenium.common.exceptions import TimeoutException
from bs4 import BeautifulSoup
import time

URL = 'https://dl.acm.org/loi/cacm/group/'

years_data_urls = []

driver = webdriver.Chrome()
driver.get(URL)
driver.implicitly_wait(2)
dec_panel = driver.find_element(By.CLASS_NAME, 'scroll')
dec_elems = dec_panel.find_elements(By.TAG_NAME, 'li')

for du in driver.find_elements(By.CLASS_NAME, 'scroll')[1:len(dec_elems)+1]:
    for year in du.find_elements(By.TAG_NAME, 'li'):
        years_data_urls.append(year.find_element(By.TAG_NAME, "a").get_attribute("data-url"))

print(years_data_urls)
# for data_url in years_data_urls:
#     driver.get(f'{URL}{years_data_urls}')

#     if len(driver.find_elements(By.ID, 'CybotCookiebotDialogBody')) != 0:
#         WebDriverWait(driver, 1).until(EC.element_to_be_clickable((By.ID, 'CybotCookiebotDialogBodyLevelButtonLevelOptinDeclineAll'))).click()
        

driver.close()