/*
 * 1251: https://leetcode.cn/problems/average-selling-price/
 */

-- Write your MySQL query statement below
select 
  p.product_id, 
  round(sum(p.price * us.units) / sum(us.units), 2) average_price 
from 
  Prices p, 
  UnitsSold us 
where 
  p.product_id = us.product_id and 
  us.purchase_date between p.start_date and p.end_date 
group by 
  p.product_id
