
-- ----------------------------
-- Sequence structure for user_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."user_id_seq";
CREATE SEQUENCE "public"."user_id_seq"
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;
ALTER SEQUENCE "public"."user_id_seq" OWNER TO "postgres";

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS "public"."user";
CREATE TABLE "public"."user" (
  "id" int8 NOT NULL DEFAULT nextval('user_id_seq'::regclass),
  "name" varchar(255) COLLATE "pg_catalog"."default",
  "name_fl" varchar(255) COLLATE "pg_catalog"."default",
  "code" varchar(80) COLLATE "pg_catalog"."default",
  "name_full" varchar(255) COLLATE "pg_catalog"."default",
  "state" int2 NOT NULL DEFAULT 1,
  "description" varchar(255) COLLATE "pg_catalog"."default",
  "create_at" timestamptz(6) DEFAULT CURRENT_TIMESTAMP,
  "update_at" timestamptz(6),
  "create_by" varchar(80) COLLATE "pg_catalog"."default" DEFAULT ''::character varying,
  "update_by" varchar(80) COLLATE "pg_catalog"."default" DEFAULT ''::character varying,
  "sort" int8 NOT NULL DEFAULT 0,
  "ano" varchar(80) COLLATE "pg_catalog"."default" DEFAULT ''::character varying
)
;
ALTER TABLE "public"."user" OWNER TO "postgres";
COMMENT ON COLUMN "public"."user"."name" IS '名称';
COMMENT ON COLUMN "public"."user"."name_fl" IS '名称外文';
COMMENT ON COLUMN "public"."user"."code" IS '编号代号';
COMMENT ON COLUMN "public"."user"."name_full" IS '全称';
COMMENT ON COLUMN "public"."user"."state" IS '1有效2停用11取消(对应有效)12弃置(对应停用)13批量删除(无状态)';
COMMENT ON COLUMN "public"."user"."description" IS '描述';
COMMENT ON COLUMN "public"."user"."create_at" IS '创建时间';
COMMENT ON COLUMN "public"."user"."update_at" IS '更新时间';
COMMENT ON COLUMN "public"."user"."create_by" IS '创建人';
COMMENT ON COLUMN "public"."user"."update_by" IS '更新人';
COMMENT ON COLUMN "public"."user"."sort" IS '排序';
COMMENT ON COLUMN "public"."user"."ano" IS '操作人';
COMMENT ON TABLE "public"."user" IS '用户';

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
SELECT setval('"public"."user_id_seq"', 4, true);

-- ----------------------------
-- Indexes structure for table user
-- ----------------------------
CREATE INDEX "idx_user_create_at" ON "public"."user" USING btree (
  "create_at" "pg_catalog"."timestamptz_ops" ASC NULLS LAST
);
CREATE INDEX "idx_user_create_by" ON "public"."user" USING btree (
  "create_by" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);
CREATE INDEX "idx_user_state" ON "public"."user" USING btree (
  "state" "pg_catalog"."int2_ops" ASC NULLS LAST
);
CREATE INDEX "user_owner_code_idx" ON "public"."user" USING btree (
  "ano" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table user
-- ----------------------------
ALTER TABLE "public"."user" ADD CONSTRAINT "user_pkey" PRIMARY KEY ("id");
