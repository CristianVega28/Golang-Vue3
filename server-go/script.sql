CREATE DATABASE vue_golang;

CREATE TABLE clientes (
    ID varchar(10),
    fullname varchar(120),
    birthdate varchar(50),
    address_client varchar(120),
    sector_client varchar(60),
    postalcode_client varchar(30),
    phone int,
    email varchar(50),
    dischargedate varchar(100),
    customergroup varchar(70)
);