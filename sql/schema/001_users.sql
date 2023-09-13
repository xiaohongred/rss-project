
-- +goose Up

CREATE TABLE users (
                       id INT NOT NULL AUTO_INCREMENT COMMENT '主键',
                       create_at TIMESTAMP NOT NULL COMMENT '创建时间',
                       update_at TIMESTAMP NOT NULL COMMENT '修改时间',
                       name VARCHAR(50) NOT NULL COMMENT '姓名',
                       PRIMARY KEY (id)
);

-- +goose Down
DROP  TABLE users;