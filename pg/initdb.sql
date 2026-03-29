create database  mydb;

create user jobuser with password 'job1234';
grant all privileges on database mydb  to jobuser;


\c mydb
\c - jobuser
create table workorder(
    id INT PRIMARY KEY NOT NULL,
    time  text  not null,
    name  text not null,
    phone text not null,
    describe  text not null,
    state int not null,
    imagepath text
);


CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    role VARCHAR(20),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);