--- create tables fo db

CREATE TABLE FEATURES(
  id serial PRIMARY KEY,
  name varchar(255) NOT NULL,
  description text 
);

CREATE TABLE TAGS(
  id serial PRIMARY KEY,
  name varchar(255) NOT NULL,
  description text 
);

CREATE TABLE BANNERS (
  id serial PRIMARY KEY,
  title varchar(255) NOT NULL,
  text text NOT NULL,
  url text NOT NULL,
  feature_id int NOT NULL,
  visible boolean NOT NULL,
  create_time timestamp NOT NULL,
  update_time timestamp NOT NULL,

  --CONSTRAINT unique_feature UNIQUE (feature_id),
  
  FOREIGN KEY (feature_id) REFERENCES FEATURES (id)
);

CREATE TABLE B_T(
  id serial PRIMARY KEY,
  banner_id int NOT NUll,
  tag_id int NOT NULL,

  CONSTRAINT unique_bid_tid UNIQUE (banner_id, tag_id),

  FOREIGN KEY (banner_id) REFERENCES BANNERS (id) ON DELETE CASCADE ON UPDATE CASCADE,
  FOREIGN KEY (tag_id) REFERENCES TAGS (id) ON DELETE CASCADE ON UPDATE CASCADE
);


