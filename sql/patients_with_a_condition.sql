/*
    1527: https://leetcode.cn/problems/patients-with-a-condition/
 */

-- Write your MySQL query statement below

-- like
select patient_id, patient_name, conditions from Patients where conditions like 'DIAB1%' or conditions like '% DIAB1%';

-- regexp
SELECT * FROM PATIENTS WHERE CONDITIONS REGEXP '^DIAB1|\\sDIAB1';
