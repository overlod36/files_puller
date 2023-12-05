from selenium import webdriver
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.common.by import By
from selenium.webdriver.support import expected_conditions as EC
from selenium.common.exceptions import TimeoutException
from bs4 import BeautifulSoup
import time

decades_id = {}
years_id = {}

driver = webdriver.Chrome()
driver.get('https://dl.acm.org/loi/cacm/group/')
# driver.implicitly_wait(10)

dec_panel = driver.find_element(By.CLASS_NAME, 'scroll')
dec_elems = dec_panel.find_elements(By.TAG_NAME, 'li')

for elem in dec_elems:
    decades_id[elem.find_element(By.TAG_NAME, 'a').get_attribute('title')] = elem.find_element(By.TAG_NAME, 'a').get_attribute('id')

dec_panel_parent = dec_panel.find_element(By.XPATH, './..')

years_panel = dec_panel_parent.find_element(By.XPATH, "following-sibling::*[1]").find_element(By.CLASS_NAME, 'scroll')
years_elems = years_panel.find_elements(By.TAG_NAME, 'li')

for elem in years_elems:
    years_id[elem.find_element(By.TAG_NAME, 'a').get_attribute('title')] = elem.find_element(By.TAG_NAME, 'a').get_attribute('id')

print(years_id)
driver.close()