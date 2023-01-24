CREATE TABLE IF NOT EXISTS urls(
   long_url varchar(255),
   short_url varchar(15),
   created_at timestamp,
   constraint urls_pk primary key (long_url, short_url)
);
