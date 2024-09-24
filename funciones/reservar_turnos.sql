create or replace function reservar_turnos(_nro_paciente int, _dni_medique int, _fecha timestamp)
returns boolean as $$
declare
	_basura record;
begin
  select * from medique m into _basura where m.dni_medique in (_dni_medique);
  if found then
    select * from paciente p into _basura where p.nro_paciente in (_nro_paciente);
    if found then
      if exists (select * from paciente p, medique m, cobertura c where p.nro_obra_social = c.nro_obra_social and c.dni_medique = m.dni_medique and p.nro_paciente = _nro_paciente and m.dni_medique = _dni_medique) then
        if exists (select * from turno t where t.fecha in (_fecha) and t.estado = 'Disponible') then
            select t.nro_paciente into _basura from turno t
            group by t.nro_paciente having count (t.nro_paciente) = 5;
            if found then
              raise notice 'supera límite de reserva de turnos.';
              return false;
            else
              update turno t
              set 
                nro_consultorio = a.nro_consultorio,
                nro_paciente = _nro_paciente,
                nro_obra_social_consulta = p.nro_obra_social,
                nro_afiliade_consulta = p.nro_afiliade,
                monto_paciente = c.monto_paciente,
                monto_obra_social = c.monto_obra_social,
                f_reserva = current_date,
                estado = 'Reservado'
              from cobertura c, paciente p, agenda a
              where
                t.dni_medique in (_dni_medique) and c.dni_medique in (_dni_medique)
                and t.fecha in (_fecha) and p.nro_paciente in (_nro_paciente)
                and a.dni_medique in (_dni_medique);
              return true;
            end if;
        else
          insert into error (f_turno, operacion, f_error, motivo)
          values (_fecha, 'Reserva', current_date, 'turno inexistente ó no disponible.');
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
      else
        insert into error (f_turno, operacion, f_error, motivo)
        values (_fecha, 'Reserva', current_date, 'obra social de paciente no atendida por le médique.');
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
    else
      insert into error (f_turno, operacion, f_error, motivo)
        values (_fecha, 'Reserva', current_date, 'nro de historia clínica no válido.');
        update error
        set nro_consultorio = subquery.nro_consultorio, 
            dni_medique = subquery.dni_medique
        from (select t.nro_turno, a.nro_consultorio, m.dni_medique
              from turno t, agenda a, medique m
              where t.dni_medique = m.dni_medique and
              t.nro_consultorio = a.nro_consultorio) as subquery
        where error.nro_error = (select (max(nro_error)) from error);
      return false;
    end if;
  else
    insert into error (f_turno, operacion, f_error, motivo)
        values (_fecha, 'Reserva', current_date, 'dni de médique no válido.');
        update error
        set nro_consultorio = subquery.nro_consultorio,
            nro_paciente = subquery.nro_paciente
        from (select t.nro_turno, a.nro_consultorio, p.nro_paciente
              from turno t, agenda a, paciente p
              where t.nro_paciente = p.nro_paciente and
              t.nro_consultorio = a.nro_consultorio) as subquery
        where error.nro_error = (select (max(nro_error)) from error);
    return false;
  end if;
end
$$ language plpgsql;