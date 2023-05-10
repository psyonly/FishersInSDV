-- "fish_type"  IS '类型 按占位符区分 [1<<k] 0:river; 1:lake; 2:ocean; 3:forest; 4:mine 5: desert; 6:sew
-- "season"     IS '季节 按占位符区分 [1<<k] 0:any; 1:spring; 2:summer; 3:autumn; 4:winter';
-- "weather"    IS '天气 按占位符区分 [1<<k] 0:any; 1:sunny; 2:rainy; 3:wendy';
INSERT INTO "public"."tbl_fish_type" 
("fish_name", "fish_code", "fish_type", "price", "season", "weather", "img_path", "weight_ratio", "create_time", "update_time") 
VALUES 
('Albacore', '0001', 1<<2, 120, (1<<3)|(1<<4), 1, 'Albacore.png', 3, 1683534410, 1683534410),
('Anchovy', '0002', 1<<2, 40, (1<<1)|(1<<3), 1, 'Anchovy.png', 20, 1683534410, 1683534410);