CREATE TABLE countries
(
    id serial primary key,
    name varchar(255) not null
);

INSERT INTO countries(name)
VALUES ('Россия');

CREATE TABLE district
(
    id serial primary key,
    name varchar(255) not null,
    country_id int references countries (id) on delete cascade not null
);

INSERT INTO district(name, country_id)
VALUES ('Южный', '1');

CREATE TABLE regions
(
    id serial primary key,
    name varchar(255) not null,
    district_id int references district (id) on delete cascade not null
);

INSERT INTO regions(name, district_id)
VALUES ('Краснодарский край', '1');

CREATE TABLE cities
(
    id serial primary key,
    name varchar(255) not null,
    region_id int references regions (id) on delete cascade not null
);

INSERT INTO cities(name, region_id)
VALUES ('Краснодар', '1');


CREATE TABLE banks
(
    id serial primary key,
    name varchar(255) not null,
    inn varchar(16) not null,
    bic varchar(16) not null,
    kpp varchar(16) not null,
    correspondentCheck varchar(16) not null,
    active boolean not null
);

INSERT INTO banks(name, inn, bic, kpp, correspondentCheck, active)
VALUES ('Тестовый банк', '0000000000000000', '0000000000000000', '0000000000000000', '0000000000000000', true);

CREATE TABLE organizations
(
    id serial primary key,
    name varchar(255) not null,
    settlementCheck varchar(16) not null,
    inn varchar(16) not null,
    bank_id int references banks (id) on delete cascade not null,
    active boolean not null
);

INSERT INTO organizations(name, settlementCheck, inn, bank_id, active)
VALUES ('Тестовый банк', '0000000000000000', '0000000000000000', '1', true);

CREATE TABLE roles
(
    id serial primary key,
    name varchar(50) not null
);

INSERT INTO roles(name)
VALUES ('E1');

CREATE TABLE salon
(
    id serial primary key,
    name varchar(255) not null,
    city_id int references cities (id) on delete cascade not null,
    organization_id int references organizations (id) on delete cascade not null,
    role_id int references roles (id) on delete cascade not null,
    active boolean not null
);

INSERT INTO salon(name, city_id, organization_id, role_id, active)
VALUES ('Волжский', '1', '1', '1', true);

CREATE TABLE users
(
    id serial primary key,
    login varchar(30) not null,
    password varchar(100) not null,
    lastname varchar(100) not null,
    firstname varchar(100) not null,
    fathername varchar(100) not null,
    email varchar(100) not null,
    active boolean not null,
    superuser boolean not null,
    staff boolean not null,
    registrationsDate timestamp with time zone,
    entryDate timestamp with time zone,
    salon_id int references salon (id) on delete cascade not null
);

CREATE TABLE rateCustomCloset
(
    id serial primary key,
    city_id int references cities (id) on delete cascade not null,
    priceCustom float not null,
    priceOptim float not null

);

CREATE TABLE rateDelivery
(
    id serial primary key,
    city_id int references cities (id) on delete cascade not null,
    price float not null,
    pricekm float not null
);

CREATE TABLE rateMeasurement
(
    id serial primary key,
    city_id int references cities (id) on delete cascade not null,
    price float not null,
    pricekm float not null,
    intervalFrom float not null,
    IntervalUntil float not null
);

CREATE TABLE TypeProduct
(
    id serial primary key,
    name varchar(255) not null
);

CREATE TABLE SeriesCloset
(
    id serial primary key,
    name varchar(255) not null,
    active boolean not null,
    typeProduct_id int references TypeProduct (id) on delete cascade not null
);

CREATE TABLE OptionSeries
(
    id serial primary key,
    name varchar(255) not null,
    active boolean not null,
    seriesCloset_id int references SeriesCloset (id) on delete cascade not null
);

CREATE TABLE SizesOption
(
    id serial primary key,
    amount_section int not null,
    height int [] not null,
    width int [] not null,
    depth int [] not null,
    optionSeries_id int references OptionSeries (id) on delete cascade not null
);

CREATE TABLE JWT_Storage
(
  id serial primary key,
  token varchar(1000) not null
);

