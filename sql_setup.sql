
drop database poll_1;
create database poll_1;
use poll_1;
create table authdb (username VARCHAR(20) UNIQUE KEY PRIMARY KEY, passwords TEXT);
create table ballot (username VARCHAR(50) PRIMARY KEY, vote_0 VARCHAR(40)); -- add as many as vote_i
INSERT into authdb (username,passwords) VALUES ('agnes', 'Hola');
INSERT into authdb (username,passwords) VALUES ('nishant', 'paper');
INSERT into authdb (username,passwords) VALUES ('hsinghc', 'singhc');

