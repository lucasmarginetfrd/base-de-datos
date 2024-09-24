create or replace function generar_liquidacion_mensual(p_nro_obra_social INT, p_desde DATE, p_hasta DATE)
returns void as $$
declare
begin
    
    insert into liquidacion_cabecera (nro_obra_social, desde, hasta, total)
    values(p_nro_obra_social, p_desde, p_hasta, 0);
    
    update liquidacion_cabecera
    set total = (select sum(t.monto_obra_social) from turno t
                 where estado = 'Atendido' 
                 and t.nro_obra_social_consulta = liquidacion_cabecera.nro_obra_social_consulta 
                 group by extract(month from t.fecha));

   update turno
    set estado = 'Liquidado'
    where fecha between p_desde and p_hasta
    and nro_obra_social_consulta = p_nro_obra_social
    and estado = 'Atendido';

    insert into liquidacion_detalle (nro_liquidacion)
    select l.nro_liquidacion
    from turno t, liquidacion_cabecera l
    where estado = 'Liquidado'
    and not exists(
      select 1
      from liquidacion_detalle d
      where d.nro_liquidacion = l.nro_liquidacion
    );
    
    update liquidacion_detalle
    set f_atencion = subquery.fecha,
        nro_afiliade = subquery.nro_afiliade,
        dni_paciente = subquery.dni_paciente,
        nombre_paciente = subquery.nombre_paciente,
        apellido_paciente = subquery.apellido_paciente,
        dni_medique = subquery.dni_medique,
        nombre_medique = subquery.nombre,
        apellido_medique = subquery.apellido,
        especialidad = subquery.especialidad,
        monto = subquery.monto_obra_social
    from (select distinct t.fecha, p.nro_afiliade, p.dni_paciente,
          p.nombre as nombre_paciente, p.apellido as apellido_paciente, m.dni_medique, m.nombre, m.apellido,
          m.especialidad, t.monto_obra_social
          from turno t, paciente p, medique m
          where t.nro_afiliade_consulta = p.nro_afiliade and
                t.dni_medique = m.dni_medique) as subquery, turno
    where turno.fecha between p_desde and p_hasta;
end
$$ language plpgsql;