## Sosmed Api (Kelompok 1)
Sosmed Api is an api that is used for social media application purposes. Meanwhile, Social Media Apps itself is an application that is used to share stories, articles or photos.

### ERD
![erd](/docs/erd_project.drawio.png)

### Api Spec
please follow [swagger hub](https://app.swaggerhub.com/apis/HERUSETIAWAN_1/sosmed/1.0.0) or this [postman workspace](https://www.postman.com/herusetiawans/workspace/alta-sosmed) to see api specifications

### How To Run
- Clone repository
    ```bash
    git clone https://github.com/Sosmed-App-Kelompok-1/BE.git SOSMED_API
    ```

- Move to cloned repository folder
    ```bash
    cd SOSMED_API
    ```

- Update dependecies
    ```bash
    go mod tidy
    ```

- Create your own database

- Copy `.env.example` to `.env`
    ```bash
    cp .env.example .env
    ```

- Change and customize `.env` file

- Run Sosmed Api
    ```bash
    go run .
    ```