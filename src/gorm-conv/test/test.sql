CREATE TABLE `account_0`(
    `version` bigint DEFAULT 0,
    `id` int DEFAULT 0,
    `account` varchar(128) NOT NULL,
    `allbinary` mediumblob,
    INDEX account_id(``),
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
CREATE TABLE `account_1`(
    `version` bigint DEFAULT 0,
    `id` int DEFAULT 0,
    `account` varchar(128) NOT NULL,
    `allbinary` mediumblob,
    INDEX account_id(``),
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
CREATE TABLE `account_2`(
    `version` bigint DEFAULT 0,
    `id` int DEFAULT 0,
    `account` varchar(128) NOT NULL,
    `allbinary` mediumblob,
    INDEX account_id(``),
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
CREATE TABLE `bag_0`(
    `version` bigint DEFAULT 0,
    `id` int DEFAULT 0,
    `allbinary` mediumblob,
    INDEX bag_id(``),
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
CREATE TABLE `bag_1`(
    `version` bigint DEFAULT 0,
    `id` int DEFAULT 0,
    `allbinary` mediumblob,
    INDEX bag_id(``),
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
CREATE TABLE `bag_2`(
    `version` bigint DEFAULT 0,
    `id` int DEFAULT 0,
    `allbinary` mediumblob,
    INDEX bag_id(``),
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
