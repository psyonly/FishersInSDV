CREATE TABLE IF NOT EXISTS "tbl_fish_type" (
		"id" int4 NOT NULL,
		"fish_name" varchar(32) NOT NULL DEFAULT '',
		"fish_code" varchar(32) NOT NULL DEFAULT '',
		"fish_type" int8 NOT NULL DEFAULT 1,
		"price" int4 NOT NULL DEFAULT 0,
		"season" int4 NOT NULL DEFAULT 1,
		"weather" int4 NOT NULL DEFAULT 1,
		"img_path" varchar(32) NOT NULL DEFAULT '',
		"weight_ratio" int4 NOT NULL DEFAULT 1,
		"create_time" int4 NOT NULL DEFAULT 0,
		"update_time" int4 NOT NULL DEFAULT 0
);
COMMENT ON COLUMN "public"."tbl_fish_type"."id" IS 'ID';
COMMENT ON COLUMN "public"."tbl_fish_type"."fish_name" IS '名称';
COMMENT ON COLUMN "public"."tbl_fish_type"."fish_code" IS '编号';
COMMENT ON COLUMN "public"."tbl_fish_type"."fish_type" IS '类型 按占位符区分 [1<<k] 0:river; 1:lake; 2:ocean; 3:forest; 4:mine 5: desert; 6:sewer; 7:void; 8:island';
COMMENT ON COLUMN "public"."tbl_fish_type"."price" IS '基础价格';
COMMENT ON COLUMN "public"."tbl_fish_type"."season" IS '季节 按占位符区分 [1<<k] 0:any; 1:spring; 2:summer; 3:autumn; 4:winter';
COMMENT ON COLUMN "public"."tbl_fish_type"."weather" IS '天气 按占位符区分 [1<<k] 0:any; 1:sunny; 2:rainy; 3:wendy';
COMMENT ON COLUMN "public"."tbl_fish_type"."img_path" IS '图片路径';
COMMENT ON COLUMN "public"."tbl_fish_type"."weight_ratio" IS '比例权重';
COMMENT ON COLUMN "public"."tbl_fish_type"."create_time" IS '创建时间';
COMMENT ON COLUMN "public"."tbl_fish_type"."update_time" IS '修改时间';

-- 表注释
COMMENT ON TABLE "public"."tbl_fish_type" IS '';
-- 更改表拥有者
ALTER TABLE "public"."tbl_fish_type" OWNER TO postgres;
-- 删除序列 强制删除依赖
DROP SEQUENCE IF EXISTS "public"."tbl_fish_type_id_seq" CASCADE;
-- 创建序列
CREATE SEQUENCE "public"."tbl_fish_type_id_seq" AS integer START WITH 1 INCREMENT BY 1 NO MINVALUE NO MAXVALUE CACHE 1;
-- 设置序列的拥有者
ALTER TABLE "public"."tbl_fish_type_id_seq" OWNER TO postgres;
-- 设置序列的拥有字段
ALTER SEQUENCE "public"."tbl_fish_type_id_seq" OWNED BY "public"."tbl_fish_type"."id";
-- 设置自增序列
ALTER TABLE ONLY "public"."tbl_fish_type" ALTER COLUMN "id" SET DEFAULT nextval('"public"."tbl_fish_type_id_seq"'::regclass);
-- 设置主键 
ALTER TABLE ONLY "public"."tbl_fish_type" ADD CONSTRAINT "public_tbl_fish_type_pkey" PRIMARY KEY ("id");
