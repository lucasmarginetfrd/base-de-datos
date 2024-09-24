create or replace function atender_turno(nro_turno_hg int) returns boolean as $$
declare
    v_estado_turno char(10);
    v_fecha_turno date;
begin
 
 select estado, fecha into v_estado_turno, v_fecha_turno
    from turno
    where  nro_turno = nro_turno_hg;
    if not found then
        insert into error (f_turno, operacion, f_error, motivo)
        values (current_date, 'Atención', current_date, 'nro de turno no válido');
        
        return false;
    end if;
    if v_estado_turno != 'Reservado' then
        insert into error (f_turno, operacion, f_error, motivo)
        values (current_date, 'Atención', current_date, 'turno no reservado');
        return false;
    end if;
    if v_fecha_turno != current_date then
        insert into error (f_turno, operacion, f_error, motivo)
        values (current_date, 'Atención', current_date, 'turno no corresponde a la fecha del día');
        update error
        set nro_consultorio = subquery.nro_consultorio, 
            dni_medique = subquery.dni_medique, 
            nro_paciente = subquery.nro_paciente
        from (select t.nro_turno, a.nro_consultorio, m.dni_medique, p.nro_paciente
              from turno t, agenda a, medique m, paciente p
              where t.nro_paciente = p.nro_paciente and 
              t.dni_medique = m.dni_medique and
              t.nro_consultorio = a.nro_consultorio) as subquery
        where error.nro_error = (select (max(nro_error)) from error);     
        return false;
    end if;
    update turno
    set estado = 'Atendido'
    where nro_turno = nro_turno_hg;
    return true;
end
$$ language plpgsql;