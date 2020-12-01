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


class SystemUnderTestProps(Enum):
    URL = "https://demo.saleor.io/"
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

    def __init__(self, is_headless = True, executable_path='/usr/bin/chromedriver'):
        self.__is_headless = is_headless
        self.__executable_path = executable_path
        self.__driver: webdriver = None

    def start_driver(self):
        chrome_options = webdriver.ChromeOptions()
        if self.__is_headless:
            chrome_options.add_argument('--headless')
            chrome_options.add_argument('--no-sandbox')
        self.__driver: webdriver = webdriver.Chrome(executable_path=self.__executable_path, options=chrome_options, service_args=['--verbose', '--log-path=/tmp/chromedriver.log'])

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
    def txt_username(self):
        return ''

    @property
    def txt_password(self):
        return ''

    @property
    def btn_login(self):
        return ''

    def __init__(self, driver: Chrome):
        self.__driver: Chrome = driver

    def navigate_to_login_page(self):
        self.__driver.get(SystemUnderTestProps.URL)

    def login_to_dashboard(self):
        self.__driver.find_element_by_id(self.txt_username).send_keys(
            SystemUnderTestProps.USER_NAME)
        self.__driver.find_element_by_id(self.txt_password).send_keys(
            SystemUnderTestProps.PASSWORD)
        self.__driver.find_element_by_id(self.btn_login).click()

    def capture_home_page_screen_shot(self):
        self.__driver.get_screenshot_as_file('home-page.png')


def run_test():
    virtual_display = create_virutal_display()
    virtual_display_control(virtual_display, VirtualActions.START)
