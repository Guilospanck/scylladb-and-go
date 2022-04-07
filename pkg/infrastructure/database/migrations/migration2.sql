CREATE KEYSPACE IF NOT EXISTS tracking WITH REPLICATION = { 'class' : 'NetworkTopologyStrategy','DC1' : 3};

USE tracking;

CREATE TABLE tracking_data (
  first_name text,
  last_name text,
  timestamp timestamp,
  location varchar,
  speed double,
  heat double,
  telepathy_powers int,
  primary key((first_name, last_name), timestamp)
) WITH CLUSTERING ORDER BY (timestamp DESC);

 INSERT INTO tracking.tracking_data ("first_name","last_name","timestamp","location","speed","heat","telepathy_powers") VALUES ('Jim','Jeffries','2017-11-11 08:05+0000','New York',1.0,3.0,17);
 INSERT INTO tracking.tracking_data ("first_name","last_name","timestamp","location","speed","heat","telepathy_powers") VALUES ('Jim','Jeffries','2017-11-11 09:05+0000','New York',2.0,4.0,27);
 INSERT INTO tracking.tracking_data ("first_name","last_name","timestamp","location","speed","heat","telepathy_powers") VALUES ('Jim','Jeffries','2017-11-11 10:05+0000','New York',3.0,5.0,37);
 INSERT INTO tracking.tracking_data ("first_name","last_name","timestamp","location","speed","heat","telepathy_powers") VALUES ('Jim','Jeffries','2017-11-11 10:22+0000','New York',4.0,12.0,47);
 INSERT INTO tracking.tracking_data ("first_name","last_name","timestamp","location","speed","heat","telepathy_powers") VALUES ('Jim','Jeffries','2017-11-11 11:05+0000','New York',4.0,9.0,87);
 INSERT INTO tracking.tracking_data ("first_name","last_name","timestamp","location","speed","heat","telepathy_powers") VALUES ('Jim','Jeffries','2017-11-11 12:05+0000','New York',4.0,24.0,57);
 INSERT INTO tracking.tracking_data ("first_name","last_name","timestamp","location","speed","heat","telepathy_powers") VALUES ('Bob','Loblaw','2017-11-11 08:05+0000','Cincinnati',2.0,6.0,5);
 INSERT INTO tracking.tracking_data ("first_name","last_name","timestamp","location","speed","heat","telepathy_powers") VALUES ('Bob','Loblaw','2017-11-11 09:05+0000','Cincinnati',4.0,1.0,10);
 INSERT INTO tracking.tracking_data ("first_name","last_name","timestamp","location","speed","heat","telepathy_powers") VALUES ('Bob','Loblaw','2017-11-11 10:05+0000','Cincinnati',6.0,1.0,15);
 INSERT INTO tracking.tracking_data ("first_name","last_name","timestamp","location","speed","heat","telepathy_powers") VALUES ('Bob','Loblaw','2017-11-11 10:22+0000','Cincinnati',8.0,3.0,6);
 INSERT INTO tracking.tracking_data ("first_name","last_name","timestamp","location","speed","heat","telepathy_powers") VALUES ('Bob','Loblaw','2017-11-11 11:05+0000','Cincinnati',10.0,2.0,3);
 INSERT INTO tracking.tracking_data ("first_name","last_name","timestamp","location","speed","heat","telepathy_powers") VALUES ('Bob','Loblaw','2017-11-11 12:05+0000','Cincinnati',12.0,10.0,60);