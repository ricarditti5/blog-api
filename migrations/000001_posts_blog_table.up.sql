CREATE TABLE posts(
id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
title varchar(150) NOT NULL,
content varchar(255),
category varchar(100),
tags TEXT[]
);

CREATE INDEX idx_posts_category ON posts(category);

CREATE INDEX idx_posts_tags ON posts USING GIN(tags);