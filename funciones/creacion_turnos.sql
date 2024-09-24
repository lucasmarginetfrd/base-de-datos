create or replace function generar_turnos_disponibles(anio int, mes int)
    returns boolean as $$
    declare
        _dias2 timestamp;
        _fecha timestamp;
        _fecha2 timestamp;
        _fecha3 timestamp;
        _fecha4 timestamp;
        _fecha5 timestamp;
        _nroturno int = 0;
        _dnis record;
        _mediquedni cursor for select distinct dni_medique from agenda order by dni_medique;
        _dias record;
        _basura record;
        _horas record;
        _hora time;
    begin
        select t.fecha into _basura from turno t
        where extract(year from t.fecha) = anio and extract(month from t.fecha) = mes;
        if found then
            return false;
        else
            if exists (select nro_turno from turno) then
                select max(nro_turno) into _nroturno from turno;
            end if;
            for _dnis in _mediquedni
            loop
                for _dias in (select dia from agenda a where a.dni_medique = _dnis.dni_medique)
                loop
                    select distinct a.hora_desde, a.hora_hasta, a.duracion_turno into _horas from agenda a where a.dni_medique = _dnis.dni_medique;
                    _fecha4 = to_timestamp(anio || '-' || mes || '-' || _dias.dia || ' ' || _horas.hora_desde, 'YYYY-MM-DD HH24:MI:SS');
                    if mes != 12 then
                        _fecha5 = to_timestamp(anio || '-' || mes + 1 || '-' || _dias.dia || ' ' || _horas.hora_hasta, 'YYYY-MM-DD HH24:MI:SS');
                    else
                        _fecha5 = to_timestamp(anio + 1 || '-' || 1 || '-' || _dias.dia || ' ' || _horas.hora_hasta, 'YYYY-MM-DD HH24:MI:SS');
                    end if;
                    for _dias2 IN SELECT generate_series(_fecha4::timestamp, _fecha5::timestamp, '7 days') as _dias2
                    loop
                            _fecha = to_timestamp(anio || '-' || mes || '-' || extract(days from _dias2) || ' ' || _horas.hora_desde, 'YYYY-MM-DD HH24:MI:SS');
                            _fecha2 = to_timestamp(anio || '-' || mes || '-' || extract(days from _dias2) || ' ' || _horas.hora_hasta, 'YYYY-MM-DD HH24:MI:SS');
                            for _hora in select generate_series(_fecha::timestamp, _fecha2::timestamp, _horas.duracion_turno) as _hora
                            loop
                                _nroturno := _nroturno + 1;
                                _fecha3 = to_timestamp(anio || '-' || mes || '-' || extract(days from _dias2) || ' ' || _hora, 'YYYY-MM-DD HH24:MI:SS');
                                insert into turno(nro_turno, fecha, dni_medique, estado) values (_nroturno, _fecha3, _dnis.dni_medique, 'Disponible');
                            end loop;
                    end loop;
                end loop;
            end loop;
            return true;
        end if;
    end
$$ language plpgsql;
