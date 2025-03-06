/*
CREATE TABLE IF NOT EXISTS article_topics (
    article_id INTEGER NOT NULL,
    topic_id INTEGER NOT NULL,
    PRIMARY KEY (article_id, topic_id),
    FOREIGN KEY (article_id) REFERENCES articles(id) ON DELETE CASCADE,
    FOREIGN KEY (topic_id) REFERENCES topics(id) ON DELETE CASCADE
);
*/
