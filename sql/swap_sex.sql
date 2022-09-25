/*
    627: https://leetcode.cn/problems/swap-salary/
 */

-- Write your MySQL query statement below

-- case when then else end
UPDATE salary
SET
    sex = CASE sex
        WHEN 'm' THEN 'f'
        ELSE 'm'
    END;

-- if
update salary set sex = if(sex = 'f', 'm', 'f');

-- ascii
update salary set sex = char(ascii('m') + ascii('f') - ascii(sex));
update salary set sex = char(211 - ascii(sex));
