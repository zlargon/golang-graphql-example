# golang-graphql-example

## Prerequisite

| Name                   | Version  | install by homebrew          |
| ---------------------- | -------- | ---------------------------- |
| git                    | >=2.9    | `brew install git`           |
| yarn                   | 1.22.x   | `brew install yarn`          |
| go                     | 1.19.4   | `brew install go`            |
| Docker Desktop and CLI | 20.10.23 | `brew install --cask docker` |
| curl                   | 7.85.0   | `brew install curl`          |
| jq                     | jq-1.6   | `brew install jq`            |

## Run database by docker compose

Docker containers are defined in `mongodb.yml`

| Service       | Port  |
| :------------ | :---- |
| MongoDB       | 27017 |
| Mongo Express | 8081  |

Run docker compose by `yarn` (scripts are defined in package.json)

```bash
yarn up    # up and run docker containers
yarn stop  # stop docker containers
yarn down  # stop and delete docker containers
yarn logs  # show logs of docker docker containers
```

## Getting started

1. Run Go server at port 8080

   ```bash
   yarn start
   ```

   | API         | Description                                                                  |
   | :---------- | :--------------------------------------------------------------------------- |
   | GET /jwt    | get jwt token with basic auth header (token = Base64("$USERNAME:$PASSWORD")) |
   | GET /       | GraphQL IDE                                                                  |
   | POST /query | graphQL endpoint with bearer auth header (token = jwt token)                 |

2. Get JWT token by http api + basic auth header

   ```bash
   # Curl Example
   # token is Base64("username:password") = dXNlcm5hbWU6cGFzc3dvcmQ=
   curl -s 'http://localhost:8080/jwt' -H 'authorization: Basic dXNlcm5hbWU6cGFzc3dvcmQ='

   # API Response
   { "token": "xxxxxx" }
   ```

3. Copy the JWT token and paste into **Headers** of the GraphQL IDE

   ```json
   {
     "Authorization": "Bearer $TOKEN"
   }
   ```

4. Create a resource by graphQL API + bearer auth header

   ```graphql
   mutation createVideo {
     createVideo(input: { title: "video 1", url: "http://hello.com", userId: "1" }) {
       id
       title
       url
       author {
         id
         name
       }
     }
   }
   ```

5. Get resources by graphQL API + bearer auth header

   ```graphql
   query getVideos {
     videos {
       id
       title
       url
       author {
         id
         name
       }
     }
   }
   ```

## Test by curl scripts

```bash
./curl/get_jwt.sh
./curl/create_video.sh
./curl/find_videos.sh
```

## Screenshots

1. Get JWT token
   ![](./screenshots/get_jwt.png)

2. GraphQL create and get resources
   ![](./screenshots/graphql.png)
