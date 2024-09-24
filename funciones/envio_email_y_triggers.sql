create or replace function enviar_email()
returns trigger as $$
declare
  dia cursor for select extract(day from current_date);
begin
  case new.estado 
  when 'Reservado' then
    insert into envio_email(f_generacion, asunto, cuerpo, f_envio, estado) 
    values (current_date, 'Reserva de turno', 'El turno fue reservado correctamente', null, 'Pendiente');
    if (date_trunc('day', current_date) != date_trunc('day', current_date - interval '1 day')) then
      if (date_trunc('day', new.fecha) = date_trunc('day', current_timestamp) + interval '2 days') then
        insert into envio_email(f_generacion, asunto, f_envio, estado)
        values (current_date, 'Recordatorio de turno', null, 'Pendiente');
        update envio_email set cuerpo = _cuerpo
        from (select t.nro_turno, t.fecha, t.dni_medique from turno t 
        where (date_trunc('day', t.fecha) = date_trunc('day', current_timestamp) + interval '2 days')) as _cuerpo
        where asunto = 'Recordatorio de turno';
      end if;
    end if;
    if exists(select t.fecha from turno t where (date_trunc('day', fecha) = date_trunc('day', current_timestamp - interval '1 day'))) then
      insert into envio_email(f_generacion, asunto, f_envio, estado)
      values (current_date, 'Perdida de turno reservado', null, 'Pendiente');
      update envio_email set cuerpo = cuerpo
        from (select t.nro_turno, t.fecha, t.dni_medique from turno t 
        where (date_trunc('day', t.fecha) = date_trunc('day', current_timestamp) - interval '1 day')) as cuerpo
        where asunto = 'Perdida de turno reservado';
    end if;
  when 'Cancelado' then
    insert into envio_email(f_generacion, asunto, cuerpo, f_envio, estado) 
    values (current_date, 'Cancelacion de turno', 'El turno fue CANCELADO', null, 'Pendiente');
  else 
    return new;
  end case;
  return new;
end
$$ language plpgsql;

create trigger email_reserva_turno
after update of estado on turno
for each row
when (new.estado = 'Reservado')
execute function enviar_email();

create trigger email_cancelar_turno
after update of estado on turno
for each row
when (new.estado = 'Cancelado')
execute function enviar_email();

create trigger email_recordatorio_turno
after update on turno
for each row
when (date_trunc('day', new.fecha) = date_trunc('day', current_timestamp) + interval '2 days')
execute function enviar_email();

create trigger email_perdida_turno
after update on turno
for each row
when (date_trunc('day', current_date) != date_trunc('day', current_date - interval '1 day'))
execute function enviar_email();