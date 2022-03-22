create table if not exists users_words (
	user_id bigserial not null references users on delete cascade,
	word_id bigserial not null references words on delete cascade,
	primary key (user_id, word_id)
);
