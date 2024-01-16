CREATE TABLE countries
(
    id serial primary key,
    name varchar(255) not null
);

CREATE TABLE dictrict
(
    id serial primary key,
    name varchar(255) not null,
    country_id int references countries (id) on delete cascade not null
);

CREATE TABLE regions
(
    id serial primary key,
    name varchar(255) not null,
    dictrict_id int references dictrict (id) on delete cascade not null
);

CREATE TABLE cities
(
    id serial primary key,
    name varchar(255) not null,
    region_id int references regions (id) on delete cascade not null
);

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

CREATE TABLE organizations
(
    id serial primary key,
    name varchar(255) not null,
    settlementCheck varchar(16) not null,
    inn varchar(16) not null,
    bank_id int references banks (id) on delete cascade not null,
    active boolean not null
);

CREATE TABLE roles
(
    id serial primary key,
    name varchar(50) not null
);

CREATE TABLE salon
(
    id serial primary key,
    name varchar(255) not null,
    city_id int references cities (id) on delete cascade not null,
    organization_id int references organizations (id) on delete cascade not null,
    role_id int references roles (id) on delete cascade not null,
    active boolean not null
);

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
    registrationsDate DATE,
    entryDate DATE,
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