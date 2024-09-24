create or replace function cancelar_turnos(
    p_dni_medique int,
    p_fecha_desde date,
    p_fecha_hasta date,
    out cantidad_cancelados int
) returns integer as $$
declare
    row_data record;
begin

    update turno
    set estado = 'Cancelado'
    where dni_medique = p_dni_medique
        and fecha between p_fecha_desde and p_fecha_hasta
        and estado in ('Disponible', 'Reservado');

    get diagnostics cantidad_cancelados = row_count;
    
    insert into reprogramacion
    select t.nro_turno, p.nombre, p.apellido, p.telefono, p.email, m.nombre, m.apellido, t.estado
    from turno t, paciente p, medique m
    where t.nro_paciente = p.nro_paciente and t.dni_medique = m.dni_medique
    and estado = 'Cancelado'
    and not exists(
      select 1
      from reprogramacion r
      where r.nro_turno = t.nro_turno
    );
    
    update reprogramacion
    set estado = 'Pendiente';
    
    return;
end;
$$ language plpgsql;