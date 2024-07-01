INSERT IGNORE INTO hello_worlds (lang, message) VALUES ('en', 'Hello World');
INSERT IGNORE INTO hello_worlds (lang, message) VALUES ('ja', 'こんにちは 世界');

-- (taro,password),(hanako,PASSWORD)
INSERT IGNORE INTO users (name, password, icon_url) VALUES ('taro', '5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8', "http://example.com/icon/taro.png");
INSERT IGNORE INTO users (name, password, icon_url) VALUES ('hanako', '0be64ae89ddd24e225434de95d501711339baeee18f009ba9b4369af27d30d60', "http://example.com/icon/hanako.png");

INSERT IGNORE INTO posts (user_id, title, body) VALUES (1, 'test1', '質問1\n改行');
INSERT IGNORE INTO posts (user_id, title, body) VALUES (1, 'test2', '質問2\n改行');

INSERT IGNORE INTO comments (user_id, post_id, body) VALUES (1, 2, 'こんにちは');
INSERT IGNORE INTO comments (user_id, post_id, body) VALUES (1, 2, '元気ですか？');
