# RS - Server
Go로 만들어 보는 내 마음대로 서버

## 사용 라이브러리
- github.com/joho/godotenv
- golang.org/x/oauth2 
- github.com/rs/cors
- github.com/lib/pq (For Postgres)

## 배포 - heroku
- 포트번호를 os.Getenv("PORT")로 잡아 주어야 한다.
- Procfile을 bin/(모듈명) 으로 해야 한다
- heroku 배포 환경에서는 .env 파일이 없고 사이트 내 환경변수로 대체하기 때문에   
.env 파일을 불러오는 부분을 주석 처리해 주어야 한다.

## DB - PostgresSQL
heroku-postgresql 서비스 이용 (limit 10000줄)   
pgAdmin4에서 테이블을 확인할 수 있다.