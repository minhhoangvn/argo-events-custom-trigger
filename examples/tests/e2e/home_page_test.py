# Licensed to the Software Freedom Conservancy (SFC) under one
# or more contributor license agreements.  See the NOTICE file
# distributed with this work for additional information
# regarding copyright ownership.  The SFC licenses this file
# to you under the Apache License, Version 2.0 (the
# "License"); you may not use this file except in compliance
# with the License.  You may obtain a copy of the License at
#
#   http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing,
# software distributed under the License is distributed on an
# "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
# KIND, either express or implied.  See the License for the
# specific language governing permissions and limitations
# under the License.
# MinhHoang
# This sample source code is used for demonstration a basic argo workflow with end2end automation test
from selenium import webdriver
from selenium.webdriver.remote.webelement import WebElement
from selenium.webdriver import Chrome
from pyvirtualdisplay import Display
from enum import Enum
from selenium.webdriver.common.by import By
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC
import time


class SystemUnderTestProps(Enum):
    URL = "https://demo.saleor.io/"
    MY_ACCOUNT_URL = "https://demo.saleor.io/account/"
    USER_NAME = "admin@example.com"
    PASSWORD = "admin"


class VirtualActions(Enum):
    START = 'start'
    STOP = 'stop'


class ChromeDriverFactory(object):
    @property
    def webdriver(self):
        if self.__driver is None:
            self.start_driver()
        return self.__driver

    def __init__(self,
                 is_headless=True,
                 executable_path='/usr/bin/chromedriver'):
        self.__is_headless = is_headless
        self.__executable_path = executable_path
        self.__driver: webdriver = None

    def start_driver(self):
        chrome_options = webdriver.ChromeOptions()
        if self.__is_headless:
            chrome_options.add_argument('--headless')
            chrome_options.add_argument('--no-sandbox')
        self.__driver: webdriver = webdriver.Chrome(
            executable_path=self.__executable_path,
            options=chrome_options,
            service_args=['--verbose', '--log-path=/tmp/chromedriver.log'])
        self.__driver.set_page_load_timeout(60)

    def quit_driver(self):
        if self.__driver is None:
            return print('Driver Session Already Terminated!!!')
        self.__driver.quit()
        self.__driver = None


def create_virutal_display(height, width):
    return Display(visible=0, size=(height, width))


def virtual_display_control(virtual_object: Display, action: str):
    if action == VirtualActions.START:
        return virtual_object.start()
    if action == VirtualActions.STOP:
        return virtual_object.stop()


class HomePage(object):
    @property
    def icon_login(self):
        return "li[data-test='desktopMenuLoginOverlayLink']"

    @property
    def txt_username(self):
        return "input[name='email']"

    @property
    def txt_password(self):
        return "input[name='password']"

    @property
    def btn_login(self):
        return "button[data-test='submit']"

    @property
    def panel_my_account(self):
        return "div[class*='account__menu'"

    def __init__(self, driver: Chrome):
        self.__driver: Chrome = driver

    def navigate_to_login_page(self):
        self.__driver.get(SystemUnderTestProps.URL.value)

    def login_to_dashboard(self):
        WebDriverWait(self.__driver, 30).until(
            EC.element_to_be_clickable((By.CSS_SELECTOR, self.icon_login)))
        self.__driver.find_element_by_css_selector(self.icon_login).click()
        WebDriverWait(self.__driver, 30).until(
            EC.element_to_be_clickable((By.CSS_SELECTOR, self.txt_username)))
        self.__driver.find_element_by_css_selector(self.txt_username).clear()
        self.__driver.find_element_by_css_selector(
            self.txt_username).send_keys(SystemUnderTestProps.USER_NAME.value)
        self.__driver.find_element_by_css_selector(self.txt_password).clear()
        self.__driver.find_element_by_css_selector(
            self.txt_password).send_keys(SystemUnderTestProps.PASSWORD.value)
        self.__driver.find_element_by_css_selector(self.btn_login).click()
        time.sleep(10)

    def capture_home_page_screen_shot(self):
        self.__driver.get_screenshot_as_file('home-page.png')

    def verify_login_successfully(self):
        self.__driver.get(SystemUnderTestProps.MY_ACCOUNT_URL.value)
        WebDriverWait(self.__driver, 30).until(
            EC.visibility_of_element_located(
                (By.CSS_SELECTOR, self.panel_my_account)))
        panel_element = self.__driver.find_element_by_css_selector(
            self.panel_my_account)
        assert panel_element is not None


def run_test():
    chrome_driver = ChromeDriverFactory(
        executable_path='/Users/minhhoang/Downloads/chromedriver',
        is_headless=False)
    chrome_driver.start_driver()
    home_page = HomePage(chrome_driver.webdriver)
    virtual_display = create_virutal_display(1920, 1080)
    virtual_display_control(virtual_display, VirtualActions.START)
    try:
        home_page.navigate_to_login_page()
        home_page.login_to_dashboard()
        home_page.verify_login_successfully()
    except Exception as e:
        print('Failed Test: ', str(e))
        raise e
    finally:
        home_page.capture_home_page_screen_shot()
        chrome_driver.quit_driver()
        virtual_display_control(virtual_display, VirtualActions.STOP)


run_test()