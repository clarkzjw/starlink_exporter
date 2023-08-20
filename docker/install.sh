sudo apt-get install curl wget build-essential htop vim git chrony telegraf -y

curl -fsSL https://get.docker.com | bash
sudo systemctl enable docker
sudo systemctl restart docker

sudo curl -SL https://github.com/docker/compose/releases/download/v2.20.3/docker-compose-linux-x86_64 -o /usr/local/bin/docker-compose && sudo chmod +x /usr/local/bin/docker-compose

sudo cp config/telegraf/telegraf.conf /etc/telegraf/telegraf.conf
sudo systemctl restart telegraf

sudo cp config/chrony/chrony.conf /etc/chrony/chrony.conf
sudo service chrony force-reload

sudo docker-compose pull
sudo docker-compose up -d
