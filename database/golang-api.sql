/*
 Navicat Premium Data Transfer

 Source Server         : Laragon
 Source Server Type    : MySQL
 Source Server Version : 80030 (8.0.30)
 Source Host           : 127.0.0.1:3306
 Source Schema         : golang-api

 Target Server Type    : MySQL
 Target Server Version : 80030 (8.0.30)
 File Encoding         : 65001

 Date: 23/07/2025 16:45:07
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for category
-- ----------------------------
DROP TABLE IF EXISTS `category`;
CREATE TABLE `category`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `phone_number` int NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of category
-- ----------------------------
INSERT INTO `category` VALUES (1, '1', NULL);
INSERT INTO `category` VALUES (2, '111', NULL);
INSERT INTO `category` VALUES (3, 'qqqq', NULL);
INSERT INTO `category` VALUES (4, 'aaaa', NULL);

-- ----------------------------
-- Table structure for data_dumy
-- ----------------------------
DROP TABLE IF EXISTS `data_dumy`;
CREATE TABLE `data_dumy`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `bulan` int NULL DEFAULT NULL,
  `tahun` int NULL DEFAULT NULL,
  `dumy` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 24 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of data_dumy
-- ----------------------------
INSERT INTO `data_dumy` VALUES (1, 1, 2020, 'test dataaaaaa');
INSERT INTO `data_dumy` VALUES (2, 1, 2020, 'test dataaaaaa');
INSERT INTO `data_dumy` VALUES (3, 2, 2020, 'test dataaaaaa');
INSERT INTO `data_dumy` VALUES (4, 4, 2020, 'test dataaaaaa');
INSERT INTO `data_dumy` VALUES (5, 5, 2020, 'test dataaaaaa');
INSERT INTO `data_dumy` VALUES (6, 1, 2021, 'test dataaaaaa');
INSERT INTO `data_dumy` VALUES (7, 2, 2021, 'test dataaaaaa');
INSERT INTO `data_dumy` VALUES (8, 3, 2021, 'test dataaaaaa');
INSERT INTO `data_dumy` VALUES (9, 5, 2021, 'test dataaaaaa');
INSERT INTO `data_dumy` VALUES (10, 1, 2022, 'test dataaaaaa');
INSERT INTO `data_dumy` VALUES (11, 2, 2022, 'test dataaaaaa');
INSERT INTO `data_dumy` VALUES (12, 2, 2022, 'test dataaaaaa');
INSERT INTO `data_dumy` VALUES (13, 4, 2022, 'test dataaaaaa');
INSERT INTO `data_dumy` VALUES (14, 5, 2022, 'test dataaaaaa');
INSERT INTO `data_dumy` VALUES (15, 2, 2022, 'test dataaaaaa');
INSERT INTO `data_dumy` VALUES (16, 3, 2022, 'test dataaaaaa');
INSERT INTO `data_dumy` VALUES (17, 4, 2022, 'test dataaaaaa');
INSERT INTO `data_dumy` VALUES (18, 5, 2022, 'test dataaaaaa');
INSERT INTO `data_dumy` VALUES (19, 1, 2024, 'test dataaaaaa');
INSERT INTO `data_dumy` VALUES (20, 2, 2024, 'test dataaaaaa');
INSERT INTO `data_dumy` VALUES (21, 2, 2024, 'test dataaaaaa');
INSERT INTO `data_dumy` VALUES (22, 2, 2024, 'test dataaaaaa');
INSERT INTO `data_dumy` VALUES (23, 2, 2024, 'test dataaaaaa');

-- ----------------------------
-- Table structure for data_excel
-- ----------------------------
DROP TABLE IF EXISTS `data_excel`;
CREATE TABLE `data_excel`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `Nama_Kolom1` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `Nama_Kolom2` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `Nama_Kolom3` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `Nama_Kolom4` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `Nama_Kolom5` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `Nama_Kolom6` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `Nama_Kolom7` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `Nama_Kolom8` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `Nama_Kolom9` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `user_id` int NULL DEFAULT NULL,
  `bulan` int NULL DEFAULT NULL,
  `tahun` int NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `data_excel_ibfk_1`(`user_id` ASC) USING BTREE,
  CONSTRAINT `data_excel_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 34720 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of data_excel
-- ----------------------------
INSERT INTO `data_excel` VALUES (34679, 'a8', 'b8', 'c8', 'd8', 'e8', 'f8', 'g8', 'h8', 'i8', 1, 1, 2020);
INSERT INTO `data_excel` VALUES (34680, 'a9', 'b9', 'c9', 'd9', 'e9', 'f9', 'g9', 'h9', 'i9', 1, 2, 2020);
INSERT INTO `data_excel` VALUES (34681, 'a10', 'b10', 'c10', 'd10', 'e10', 'f10', 'g10', 'h10', 'i10', 1, 3, 2020);
INSERT INTO `data_excel` VALUES (34682, 'a11', 'b11', 'c11', 'd11', 'e11', 'f11', 'g11', 'h11', 'i11', 1, 4, 2020);
INSERT INTO `data_excel` VALUES (34683, 'a12', 'b12', 'c12', 'd12', 'e12', 'f12', 'g12', 'h12', 'i12', 1, 5, 2020);
INSERT INTO `data_excel` VALUES (34684, 'a1', 'b1', 'c1', 'd1', 'e1', 'f1', 'g1', 'h1', 'i1', 1, 1, 2021);
INSERT INTO `data_excel` VALUES (34685, 'a2', 'b2', 'c2', 'd2', 'e2', 'f2', 'g2', 'h2', 'i2', 1, 2, 2021);
INSERT INTO `data_excel` VALUES (34686, 'a3', 'b3', 'c3', 'd3', 'e3', 'f3', 'g3', 'h3', 'i3', 1, 3, 2021);
INSERT INTO `data_excel` VALUES (34687, 'a4', 'b4', 'c4', 'd4', 'e4', 'f4', 'g4', 'h4', 'i4', 1, 4, 2021);
INSERT INTO `data_excel` VALUES (34688, 'a5', 'b5', 'c5', 'd5', 'e5', 'f5', 'g5', 'h5', 'i5', 1, 1, 2022);
INSERT INTO `data_excel` VALUES (34689, 'a6', 'b6', 'c6', 'd6', 'e6', 'f6', 'g6', 'h6', 'i6', 1, 2, 2022);
INSERT INTO `data_excel` VALUES (34690, 'a7', 'b7', 'c7', 'd7', 'e7', 'f7', 'g7', 'h7', 'i7', 1, 3, 2022);
INSERT INTO `data_excel` VALUES (34707, 'a12', 'b12', 'c12', 'd12', 'e12', 'f12', 'g12', 'h12', 'i12', 1, 4, 2022);
INSERT INTO `data_excel` VALUES (34708, 'a1', 'b1', 'c1', 'd1', 'e1', 'f1', 'g1', 'h1', 'i1', 2, 5, 2022);
INSERT INTO `data_excel` VALUES (34709, 'a2', 'b2', 'c2', 'd2', 'e2', 'f2', 'g2', 'h2', 'i2', 2, 1, 2023);
INSERT INTO `data_excel` VALUES (34710, 'a3', 'b3', 'c3', 'd3', 'e3', 'f3', 'g3', 'h3', 'i3', 2, 2, 2023);
INSERT INTO `data_excel` VALUES (34711, 'a4', 'b4', 'c4', 'd4', 'e4', 'f4', 'g4', 'h4', 'i4', 2, 3, 2023);
INSERT INTO `data_excel` VALUES (34712, 'a5', 'b5', 'c5', 'd5', 'e5', 'f5', 'g5', 'h5', 'i5', 2, 4, 2023);
INSERT INTO `data_excel` VALUES (34713, 'a6', 'b6', 'c6', 'd6', 'e6', 'f6', 'g6', 'h6', 'i6', 2, 5, 2023);
INSERT INTO `data_excel` VALUES (34714, 'a7', 'b7', 'c7', 'd7', 'e7', 'f7', 'g7', 'h7', 'i7', 2, 1, 2024);
INSERT INTO `data_excel` VALUES (34715, 'a8', 'b8', 'c8', 'd8', 'e8', 'f8', 'g8', 'h8', 'i8', 2, 2, 2024);
INSERT INTO `data_excel` VALUES (34716, 'a9', 'b9', 'c9', 'd9', 'e9', 'f9', 'g9', 'h9', 'i9', 2, 3, 2024);
INSERT INTO `data_excel` VALUES (34718, 'a11', 'b11', 'c11', 'd11', 'e11', 'f11', 'g11', 'h11', 'i11', 2, 4, 2024);
INSERT INTO `data_excel` VALUES (34719, 'a12', 'b12', 'c12', 'd12', 'e12', 'f12', 'g12', 'h12', 'i12', 2, 5, 2024);

-- ----------------------------
-- Table structure for import_status
-- ----------------------------
DROP TABLE IF EXISTS `import_status`;
CREATE TABLE `import_status`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `import_file_path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `import_status` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `import_start` int NULL DEFAULT NULL,
  `import_batch` int NULL DEFAULT NULL,
  `import_total_row` int NULL DEFAULT NULL,
  `User_id` int NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 18 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of import_status
-- ----------------------------
INSERT INTO `import_status` VALUES (1, 'RjpcRE9LVU1FTlRBU0kgS09ESU5HXEdvbGFuZ1xnb2xhbmdfYXBpXHB1YmxpY1x1cGxvYWRzXGRhdGFfMTc1MzE3NTQzNS54bHN4', 'completed', 33657, 1000, 33657, 0);
INSERT INTO `import_status` VALUES (7, 'RjpcRE9LVU1FTlRBU0kgS09ESU5HXEdvbGFuZ1xnb2xhbmdfYXBpXHB1YmxpY1x1cGxvYWRzXGRhdGFfMTc1MzE3MzE5NS54bHN4', 'completed', 20, 5, 20, 0);
INSERT INTO `import_status` VALUES (9, 'RjpcRE9LVU1FTlRBU0kgS09ESU5HXEdvbGFuZ1xnb2xhbmdfYXBpXHB1YmxpY1x1cGxvYWRzXGRhdGFfMTc1MzE3NTQxMC54bHN4', 'processing', 1, 5, 33657, 0);
INSERT INTO `import_status` VALUES (11, 'RjpcRE9LVU1FTlRBU0kgS09ESU5HXEdvbGFuZ1xnb2xhbmdfYXBpXHB1YmxpY1x1cGxvYWRzXGRhdGFfMTc1MzI0NTE2Ny54bHN4', 'processing', 1001, 1000, 33657, 0);
INSERT INTO `import_status` VALUES (12, 'RjpcRE9LVU1FTlRBU0kgS09ESU5HXEdvbGFuZ1xnb2xhbmdfYXBpXHB1YmxpY1x1cGxvYWRzXGRhdGFfMTc1MzI1MTUyNy54bHN4', 'processing', 2, 1000, 2, 1);
INSERT INTO `import_status` VALUES (13, 'RjpcRE9LVU1FTlRBU0kgS09ESU5HXEdvbGFuZ1xnb2xhbmdfYXBpXHB1YmxpY1x1cGxvYWRzXGRhdGFfMTc1MzI1MTYwOC54bHN4', 'processing', 12, 1000, 12, 2);
INSERT INTO `import_status` VALUES (14, 'RjpcRE9LVU1FTlRBU0kgS09ESU5HXEdvbGFuZ1xnb2xhbmdfYXBpXHB1YmxpY1x1cGxvYWRzXGRhdGFfMTc1MzI1MTY1Ny54bHN4', 'completed', 12, 1000, 12, 2);
INSERT INTO `import_status` VALUES (15, 'RjpcRE9LVU1FTlRBU0kgS09ESU5HXEdvbGFuZ1xnb2xhbmdfYXBpXHB1YmxpY1x1cGxvYWRzXGRhdGFfMTc1MzI1MTc1Ni54bHN4', 'processing', 12, 1000, 12, 0);
INSERT INTO `import_status` VALUES (16, 'RjpcRE9LVU1FTlRBU0kgS09ESU5HXEdvbGFuZ1xnb2xhbmdfYXBpXHB1YmxpY1x1cGxvYWRzXGRhdGFfMTc1MzI1MTc5My54bHN4', 'processing', 12, 1000, 12, 0);
INSERT INTO `import_status` VALUES (17, 'RjpcRE9LVU1FTlRBU0kgS09ESU5HXEdvbGFuZ1xnb2xhbmdfYXBpXHB1YmxpY1x1cGxvYWRzXGRhdGFfMTc1MzI1MTk5Ni54bHN4', 'completed', 12, 1000, 12, 0);

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `username` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `password` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `user_key` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `refresh_token` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO `users` VALUES (1, 'user123', '$2a$10$HDd2EH2k4VF12zv17Q263eIG92F18RaCDT5jH.oVEFTA3h/srKHr6', '2h2a0p8y1b9i9t9day', 'F1uowHRFD9oDl0RxF+BKusyZwM9J6DWUrWCuyfnn9y/LTywI4RTI2fgQdRBw45k9Gl5dueOYfsMpuZ2S35PqOLzl3eV9nAxcqgMIMo3I');
INSERT INTO `users` VALUES (2, 'user2', '$2a$10$KyJXbqlt.qZIafEMPgUe8uM9tIx9oe2Ui9J8UFylLYO0tStR9E8pG', '2h2a0p8y1b9i9t9diy', 'qLQV7E3E9QHDFBUUPLJgTnwDoV/8FcFwl+2Aki6VoC+7DHKXy+zM8a3StXKt1PVVdi358ezh0yNmJyrtjbxMBbeYPIE+Z2hXeqSXpF1e');

SET FOREIGN_KEY_CHECKS = 1;
