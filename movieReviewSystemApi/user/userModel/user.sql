CREATE TABLE users (
   user_id BIGINT AUTO_INCREMENT,
   phone VARCHAR(255)  NOT NULL UNIQUE ,
   email VARCHAR(255) ,
   password VARCHAR(255) NOT NULL,
   nick_name VARCHAR(255),
   avatar VARCHAR(255),
   gender VARCHAR(10),
   birth_date BIGINT,
   role VARCHAR(50),
   status BIGINT,
    create_date BIGINT,
    updata_date BIGINT,
   PRIMARY KEY (user_id)
);