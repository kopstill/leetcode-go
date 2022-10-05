/*
    1587: https://leetcode.cn/problems/bank-account-summary-ii/
 */

-- Write your MySQL query statement below

select 
    u.name, sum(t.amount) balance 
from Users u 
    left join Transactions t 
        on u.account = t.account 
group by u.account 
having balance > 10000;
