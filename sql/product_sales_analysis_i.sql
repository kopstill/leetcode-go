/*
    1068: https://leetcode.cn/problems/product-sales-analysis-i/
 */

-- Write your MySQL query statement below

-- left join
select p.product_name, s.year, s.price from sales s left join product p on s.product_id = p.product_id

-- inner join
select p.product_name, s.year, s.price from sales s, product p on s.product_id = p.product_id
