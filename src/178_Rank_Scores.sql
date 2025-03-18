Create table If Not Exists Scores (id int, score DECIMAL(3,2));
    Truncate table Scores;
    insert into Scores (id, score) values ('1', '3.5');
    insert into Scores (id, score) values ('2', '3.65');
    insert into Scores (id, score) values ('3', '4.0');
    insert into Scores (id, score) values ('4', '3.85');
    insert into Scores (id, score) values ('5', '4.0');
    insert into Scores (id, score) values ('6', '3.65');

with cte as
         (select score
               , ROW_NUMBER() over (order by score desc ) as rn
          from scores
          group by score)

select cte.score, cte.rn as `rank`
from scores
left join cte on scores.score = cte.score
order by score desc;