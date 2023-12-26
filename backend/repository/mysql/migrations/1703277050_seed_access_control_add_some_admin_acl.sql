-- +migrate Up
INSERT INTO `permissions`(`id`, `title`) values(3,'laptop-list');
INSERT INTO `permissions`(`id`, `title`) values(4,'laptop-delete');
INSERT INTO `permissions`(`id`, `title`) values(5,'laptop-update');

INSERT INTO `access_controls`(`actor_type`, `actor_id`, `permission_id`) values('role',2,3);
INSERT INTO `access_controls`(`actor_type`, `actor_id`, `permission_id`) values('role',2,4);
INSERT INTO `access_controls`(`actor_type`, `actor_id`, `permission_id`) values('role',2,5);

-- +migrate Down
DELETE FROM `permissions` WHERE `id` IN (3,4,5);
