DROP TABLE IF EXISTS "virtual_machines";
CREATE TABLE "virtual_machines" (
  "id" SERIAL PRIMARY KEY,
  "name" VARCHAR(64) NOT NULL UNIQUE,
  "cpu_count" INT NOT NULL,
  "connected_discs" VARCHAR(256)
);

DROP TABLE IF EXISTS "discs";
CREATE TABLE "discs" (
  "id" SERIAL PRIMARY KEY,
  "name" VARCHAR(50) NOT NULL,
  "disk_space" bigint NOT NULL
);

INSERT INTO discs (id, name, disk_space) VALUES (0, 'Intel', 4294967296);
INSERT INTO discs (id, name, disk_space) VALUES (1, 'Toshiba', 8589934592);
INSERT INTO discs (id, name, disk_space) VALUES (2, 'Toshiba2', 17179869184);

INSERT INTO "virtual_machines" (id, name, cpu_count, connected_discs) VALUES (0, 'vm0', 4, '0, 1');
INSERT INTO "virtual_machines" (id, name, cpu_count, connected_discs) VALUES (1, 'vm1', 8, '2');
