DROP TABLE IF EXISTS forums;
CREATE TABLE forums (
  "id" SERIAL PRIMARY KEY,
  "name" VARCHAR(64) NOT NULL UNIQUE,
  "topic_keyword" VARCHAR(64),
  "subscribed_users" VARCHAR(256)
);

INSERT INTO forums (id, name, topic_keyword, subscribed_users) 
  VALUES (0, 'How to create library for c++ by using golang', 'programming', 'marina');

INSERT INTO forums (id, name, topic_keyword, subscribed_users) 
  VALUES (1, 'Birthday of Yakunovich', 'politics', '');
  
INSERT INTO forums (id, name, topic_keyword, subscribed_users) 
  VALUES (2, 'Best sounds 2020', 'music', '');

