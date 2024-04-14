-- insert test_data to tables

-- Вставка данных в таблицу FEATURES
INSERT INTO FEATURES(name, description) VALUES 
  ('Feature 1', 'Description of Feature 1'),
  ('Feature 2', 'Description of Feature 2'),
  ('Feature 3', 'Description of Feature 3'),
  ('Feature 4', 'Free feature, for patch/post banner 4');

-- Вставка данных в таблицу TAGS
INSERT INTO TAGS(name, description) VALUES 
  ('Tag 1', 'Description of Tag 1'),
  ('Tag 2', 'Description of Tag 2'),
  ('Tag 3', 'Description of Tag 3');

-- Вставка данных в таблицу BANNERS
INSERT INTO BANNERS(title ,text ,url ,feature_id,visible,create_time,update_time) VALUES 
  ('Banner 1','some text for banner 1 etc..','http://url_for_banner1', 1,true,NOW(),NOW()),
  ('Banner 2','some text for banner2 etc..','http://url_for_banner2', 2,false,NOW(),NOW()),
  ('Banner 3','some text for banner 3 ','http://url_for_banner3', 3,true,NOW(),NOW());

-- Вставка данных в таблицу B_T
INSERT INTO B_T(banner_id, tag_id) VALUES 
  (1, 1),
  (1, 2),
  (2, 2),
  (3, 1),
  (3, 3),
  (1, 3),
  (2, 1);