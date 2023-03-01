## Install docker
The goal is to install docker on an instance and start application using ansible.

# 🚀 Getting Started
- Allow HTTP traffic on your instance
- Update server details in [hosts.ini](inventory/hosts.ini)
- Change directory to `configs`
    ```shell
    $ cd configs
    ```
- Run [docker](playbooks/docker.yml) playbook to configure docker on server

    ```shell
    $ ansible-playbook playbooks/docker.yml
    ```
- Run [todo](playbooks/todo.yml) playbook to start app on server

    ```shell
    $ ansible-playbook playbooks/todo.yml
    ```
- Open `http://<server_public_ip_address>` in your favorite browser

# 🛠️ Troubleshooting

- CORS policy related
  - For chrome, disable `Block insecure private network requests` flag, to do so head over to `chrome://flags/#block-insecure-private-network-requests`
- Ansible `todo` playbook
  ```shell
  msg: 'Error connecting: Error while fetching server API version: (''Connection aborted.'', PermissionError(13, ''Permission denied''))'
  ```
  wait 2-3 minutes and try again.
