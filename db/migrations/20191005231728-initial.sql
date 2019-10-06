
-- +migrate Up
CREATE TABLE articles (
	id serial NOT NULL,
	title varchar,
	CONSTRAINT article_pk PRIMARY KEY (id)
);

CREATE TABLE comments (
	id serial NOT NULL,
	article_id integer NOT NULL,
	content text,
	CONSTRAINT comment_pk PRIMARY KEY (id)
);

ALTER TABLE comments ADD CONSTRAINT article_fk FOREIGN KEY (article_id)
REFERENCES articles (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;


-- +migrate Down
ALTER TABLE comments DROP CONSTRAINT IF EXISTS article_fk CASCADE;
DROP TABLE IF EXISTS comments CASCADE;
DROP TABLE IF EXISTS articles CASCADE;
