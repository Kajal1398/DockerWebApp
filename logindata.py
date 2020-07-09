from selenium import webdriver
import time
from selenium.webdriver.common.keys import Keys
from selenium.webdriver.common.action_chains import ActionChains


options = webdriver.ChromeOptions()
options.add_argument("--start-maximized")
driver = webdriver.Chrome('C:\\Users\\Kajal\\PycharmProjects\\SeleniumTest1\\drivers\\chromedriver.exe')
action = ActionChains(driver)
time.sleep(1)

main=driver.get('https://mysterious-shelf-22386.herokuapp.com/')
time.sleep(1)

driver.maximize_window()

login=driver.find_element_by_xpath('/html/body/center/div/center/form/input[1]')
login.send_keys('kajalskankariya@gmail.com')
time.sleep(3)

pwd=driver.find_element_by_xpath('/html/body/center/div/center/form/input[2]')
pwd.send_keys('123')
time.sleep(3)

submit=driver.find_element_by_xpath('/html/body/center/div/center/form/input[3]')
submit.click()
time.sleep(3)

more=driver.find_element_by_xpath('/html/body/form/center/table/tbody/tr[1]/td[2]/input')
more.click()
more.click()

time.sleep(2)