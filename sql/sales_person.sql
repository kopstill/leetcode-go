/*
    607: https://leetcode.cn/problems/sales-person/
 */

-- Write your MySQL query statement below

select name from SalesPerson sp 
    where sp.sales_id not in (
        select o.sales_id from Orders o 
            where o.com_id = (
                select c.com_id from Company c where c.name = 'RED'
            )
    );

SELECT
    s.name
FROM
    salesperson s
WHERE
    s.sales_id NOT IN (SELECT
            o.sales_id
        FROM
            orders o
                LEFT JOIN
            company c ON o.com_id = c.com_id
        WHERE
            c.name = 'RED')
;
