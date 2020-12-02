# This dockerfile is used for demonstration a basic argo workflow with end2end automation test
# We should not run container with a root privilege 
FROM python:3.8.6

ARG chromedriver_version=87.0.4280.20

RUN APT_KEY_DONT_WARN_ON_DANGEROUS_USAGE=1 \
  && curl -sS -o - https://dl-ssl.google.com/linux/linux_signing_key.pub | apt-key add - \
  && echo 'deb [arch=amd64]  http://dl.google.com/linux/chrome/deb/ stable main' >> /etc/apt/sources.list.d/google-chrome.list \
  && apt-get update \
  && apt list --upgradable \
  && apt-get upgrade -y \
  && apt-get install -y --no-install-recommends wget \
  google-chrome-stable  unzip xvfb libxi6 libgconf-2-4 ssh sshpass nginx curl sudo \
  && mkdir /chrome-headless/ \
  && wget https://chromedriver.storage.googleapis.com/${chromedriver_version}/chromedriver_linux64.zip --directory-prefix=/chrome-headless \
  && unzip /chrome-headless/chromedriver_linux64.zip -d /chrome-headless \
  && mv /chrome-headless/chromedriver /usr/bin/chromedriver \
  && chown root:root /usr/bin/chromedriver\
  && echo 'PasswordAuthentication yes' >> /etc/ssh/sshd_config \
  && echo 'PermitRootLogin yes' >> /etc/ssh/sshd_config \
  && echo "root:root" | chpasswd \
  && service ssh start \
  && service nginx start \
  && rm -rf /var/lib/apt/lists/* 

EXPOSE 80 443 22

ENTRYPOINT ["bash", "-c", "service nginx start && service ssh start && tail -f /dev/null"]