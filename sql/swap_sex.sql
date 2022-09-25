/*
    627: https://leetcode.cn/problems/swap-salary/
 */

-- Write your MySQL query statement below

UPDATE salary
SET
    sex = CASE sex
        WHEN 'm' THEN 'f'
        ELSE 'm'
    END;

update salary set sex = if(sex = 'f', 'm', 'f');
