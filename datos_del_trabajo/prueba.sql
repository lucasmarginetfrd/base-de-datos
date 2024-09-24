select * from (select * row_number() over(partition by dni_medique order by dni_medique) as row_num from turno) subquery where row_num <= 3;
