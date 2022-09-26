/*
    607: https://leetcode.cn/problems/sales-person/
 */

-- Write your MySQL query statement below

select name from SalesPerson sp 
    where sales_id not in (
        select sales_id from Orders o 
            where o.com_id = (
                select com_id from Company where name = 'RED'
            )
    );
