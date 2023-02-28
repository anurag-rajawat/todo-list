## Install docker
The goal is to install docker on an instance and start application using ansible.

# 🚀 Getting Started
- Update server details in [hosts.ini](inventory/hosts.ini)
- Run [docker](playbooks/docker.yml) playbook to configure docker on server

    ```shell
    $ ansible-playbook configs/playbooks/docker.yml
    ```
- Run [todo](playbooks/todo.yml) playbook to start app on server

    ```shell
    $ ansible-playbook configs/playbooks/todo.yml
    ```
- Open `http://<server_public_ip_address>` in your favorite browser
