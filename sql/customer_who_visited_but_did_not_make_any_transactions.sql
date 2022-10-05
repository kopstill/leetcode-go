/*
    1581: https://leetcode.cn/problems/customer-who-visited-but-did-not-make-any-transactions/
 */

-- Write your MySQL query statement below

select 
    customer_id, count(*) as count_no_trans 
from 
    Visits 
where visit_id not in (
    select visit_id from Transactions
) 
group by customer_id;

SELECT 
    customer_id, COUNT(customer_id) count_no_trans
FROM 
    visits v
LEFT JOIN 
    transactions t ON v.visit_id = t.visit_id
-- WHERE amount IS NULL
WHERE transaction_id IS NULL
GROUP BY customer_id;
