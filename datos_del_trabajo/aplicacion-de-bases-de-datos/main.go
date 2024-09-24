package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"
	_ "github.com/lib/pq"
)

func pausa() {
	fmt.Println("Presiona Enter para continuar...")
	fmt.Scanln()
	fmt.Println("Continuando...")
}
func main() {
	db1, err := sql.Open("postgres", "user=postgres password=123 host=localhost dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db1.Close()
	db2, err := sql.Open("postgres", "user=postgres password=123 host=localhost dbname=centro_medico sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db2.Close()
	var opcion int
	for {
		fmt.Printf("Que quieres hacer ? ingresa la tecla indicada\n-Eliminar base de datos : 1\n-Crear base de datos : 2\n-Crear tablas : 3\n-Agregar keys : 4\n-Eliminar keys : 5\n-Insertar datos : 6\n-Funciones y triggers : 7\n-Salir : 0\n")
		fmt.Scan(&opcion)
		if 1 == opcion {
			dropDatabase(db1)
		}
		if 2 == opcion {
			createDatabase(db1)
		}
		if 3 == opcion {
			crearTablas(db2)
		}
		if 4 == opcion {
			agregarKeys(db2)
		}
		if 5 == opcion {
			borrarKeys(db2)
		}
		if 6 == opcion {
			insertarDatos(db2)
		}
		if 7 == opcion {
			trigger_func(db2)
		}
		if 0 == opcion {
			fmt.Printf("Saliste del programa..\n")
			break
		}
	}
}
func trigger_func(db *sql.DB){
	var opcion int
	for{
		fmt.Printf("\nMenu trigers\n-Generacion de turnos disponibles : 1\n-Generar turnos disponibles : 2\n-Reserva de turno : 3\n-Cancelacion de turnos : 4\n-Atencion de turnos : 5\n-Liquidacion para obras sociales : 6\n-Envio de emails a pacientes : 7\n-Volver : 0\n")
		fmt.Scan(&opcion)
		if 1 == opcion {
			generarTurnos(db)
		}
		if 2 == opcion {
			crearTurnos(db)
		}
		if 3 == opcion {
			reservarFuncion(db)
			reservarTurnos(db)
		}
		if 4 == opcion {
			//anularFuncion(db)
			//anularTurnos(db)
		}
		if 5 == opcion {
			//atencionFuncion(db)
			//atencionTurnos(db)
		}
		if 6 == opcion {
			//liqObraSocialFuncion(db)
			//generarLiqObraSocial(db)
		}
		if 7 == opcion {
			//envioEmail(db)
		}
		if 0 == opcion {
			fmt.Printf("Volviendo..")
			break
		}
	}
}

func dropDatabase(db *sql.DB) {
	_, err := db.Exec(`drop database centro_medico`)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Se elimino la bases de datos\n")
}
func createDatabase(db *sql.DB) {
	_, err := db.Exec(`create database centro_medico`)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Se creo la base de datos centro_medico\n")
}
func crearTablas(db *sql.DB) {
	_, err := db.Exec(`create table paciente(
                        nro_paciente int,
                        nombre text,
                        apellido text,
                        dni_paciente int,
                        f_nac date,
                        nro_obra_social int,
                        nro_afiliade int,
                        domicilio text,
                        telefono char(12),
                        email text
                    );
                    create table medique(
                        dni_medique int,
                        nombre text,
                        apellido text,
                        especialidad varchar(64),
                        monto_consulta_privada decimal(12, 2),
                        telefono char(12)
                    );create table consultorio(
                        nro_consultorio int,
                        nombre text,
                        domicilio text,
                        codigo_postal char(8),
                        telefono char(12)
                    );

                    create table agenda(
                        dni_medique int,
                        dia int,
                        nro_consultorio int,
                        hora_desde time,
                        hora_hasta time,
                        duracion_turno interval
                    );

                    create table turno(
                        nro_turno int,
                        fecha timestamp,
                        nro_consultorio int,
                        dni_medique int,
                        nro_paciente int,
                        nro_obra_social_consulta int,
                        nro_afiliade_consulta int,
                        monto_paciente decimal(12, 2),
                        monto_obra_social decimal(12, 2),
                        f_reserva timestamp,
                        estado char(10)
                    );

                    create table reprogramacion(
                        nro_turno int,
                        nombre_paciente text,
                        apellido_paciente text,
                        telefono_paciente char(12),
                        email_paciente text,
                        nombre_medique text,
                        apellido_medique text,
                        estado char(12)
                    );

                    create table error(
                        nro_error int,
                        f_turno timestamp,
                        nro_consultorio int,
                        dni_medique int,
                        nro_paciente int,
                        operacion char(12),
                        f_error timestamp,
                        motivo varchar(64)
                    );

                    create table cobertura(
                        dni_medique int,
                        nro_obra_social int,
                        monto_paciente decimal(12, 2),
                        monto_obra_social decimal(12, 2)
                    );

                    create table obra_social (
                        nro_obra_social int,
                        nombre text,
                        contacto_nombre text,
                        contacto_apellido text,
                        contacto_telefono char(12),
                        contacto_email text
                    );

                    create table liquidacion_cabecera(
                        nro_liquidacion int,
                        nro_obra_social int,
                        desde date,
                        hasta date,
                        total decimal(15, 2)
                    );

                    create table liquidacion_detalle(
                        nro_liquidacion int,
                        nro_linea int,
                        f_atencion date,
                        nro_afiliade int,
                        dni_paciente int,
                        nombre_paciente text,
                        apellido_paciente text,
                        dni_medique int,
                        nombre_medique text,
                        apellido_medique text,
                        especialidad varchar(64),
                        monto decimal(12, 2)
                    );

                    create table envio_email(
                        nro_email int,
                        f_generacion timestamp,
                        email_paciente text,
                        asunto text,
                        cuerpo text,
                        f_envio timestamp,
                        estado char(10)
                    );

                    create table solicitud_reservas(
                        nro_orden int,
                        nro_paciente int,
                        dni_medique int,
                        fecha date,
                        hora time
                    );`)
	fmt.Printf("Se agregaron las tablas\n")
	if err != nil {
		log.Fatal(err)
	}
}

func agregarKeys(db *sql.DB) {
	_, err := db.Exec(`alter table paciente add constraint pk_paciente primary key(nro_paciente);
						alter table medique add constraint pk_medique primary key(dni_medique);
						alter table consultorio add constraint pk_consultorio primary key(nro_consultorio);
						alter table turno add constraint pk_turno primary key(nro_turno);
						alter table error add constraint pk_error primary key(nro_error);
						alter table obra_social add constraint pk_obra_social primary key(nro_obra_social);
						alter table liquidacion_cabecera add constraint pk_liquidacion_cabecera primary key(nro_liquidacion);
						alter table liquidacion_detalle add constraint pk_liquidacion_detalle primary key(nro_linea);
						alter table envio_email add constraint pk_envio_email primary key(nro_email);
						alter table agenda add constraint fk_agenda_dni_medique foreign key(dni_medique) references medique(dni_medique);
						alter table agenda add constraint fk_agenda_nro_consultorio foreign key(nro_consultorio) references consultorio(nro_consultorio);
						alter table turno add constraint fk_turno_dni_medique foreign key(dni_medique) references medique(dni_medique);
						alter table turno add constraint fk_turno_nro_paciente foreign key(nro_paciente) references paciente(nro_paciente);
						alter table turno add constraint fk_turno_nro_consultorio foreign key(nro_consultorio) references consultorio(nro_consultorio);
						alter table reprogramacion add constraint fk_reprogramacion_nro_turno foreign key(nro_turno) references turno(nro_turno);
						alter table error add constraint fk_error_dni_medique foreign key(dni_medique) references medique(dni_medique);
						alter table error add constraint fk_error_nro_paciente foreign key(nro_paciente) references paciente(nro_paciente);
						alter table error add constraint fk_error_nro_consultorio foreign key(nro_consultorio) references consultorio(nro_consultorio);
						alter table cobertura add constraint fk_cobertura_dni_medique foreign key(dni_medique) references medique(dni_medique);
						alter table cobertura add constraint fk_cobertura_nro_obra_social foreign key(nro_obra_social) references obra_social(nro_obra_social);
						alter table liquidacion_cabecera add constraint fk_liquidacion_c_nro_obra_social foreign key(nro_obra_social) references obra_social(nro_obra_social);
						alter table liquidacion_detalle add constraint fk_liquidacion_d_liquidacion_c foreign key(nro_liquidacion) references liquidacion_cabecera(nro_liquidacion);
						alter table solicitud_reservas add constraint fk_solicitud_reservas_nro_paciente foreign key(nro_paciente) references paciente(nro_paciente);
						alter table solicitud_reservas add constraint fk_solicitud_reservas_dni_medique foreign key(dni_medique) references medique(dni_medique);
	`)
	fmt.Printf("Se insertaron las KEYS\n")
	if err != nil {
		log.Fatal(err)
	}
}

func borrarKeys(db *sql.DB) {
	_, err := db.Exec(`alter table agenda drop constraint fk_agenda_dni_medique; 
						alter table agenda drop constraint fk_agenda_nro_consultorio; 
						alter table turno drop constraint fk_turno_dni_medique; 
						alter table turno drop constraint fk_turno_nro_paciente; 
						alter table turno drop constraint fk_turno_nro_consultorio; 
						alter table reprogramacion drop constraint fk_reprogramacion_nro_turno; 
						alter table error drop constraint fk_error_dni_medique; 
						alter table error drop constraint fk_error_nro_paciente; 
						alter table error drop constraint fk_error_nro_consultorio; 
						alter table cobertura drop constraint fk_cobertura_dni_medique; 
						alter table cobertura drop constraint fk_cobertura_nro_obra_social; 
						alter table liquidacion_cabecera drop constraint fk_liquidacion_c_nro_obra_social; 
						alter table liquidacion_detalle drop constraint fk_liquidacion_d_liquidacion_c; 
						alter table solicitud_reservas drop constraint fk_solicitud_reservas_nro_paciente;
						alter table solicitud_reservas drop constraint fk_solicitud_reservas_dni_medique;
						alter table paciente drop constraint pk_paciente; 
						alter table medique drop constraint pk_medique;
						alter table consultorio drop constraint pk_consultorio;  
						alter table turno drop constraint pk_turno; 
						alter table error drop constraint pk_error; 
						alter table obra_social drop constraint pk_obra_social; 
						alter table liquidacion_cabecera drop constraint pk_liquidacion_cabecera; 
						alter table liquidacion_detalle drop constraint pk_liquidacion_detalle; 
						alter table envio_email drop constraint pk_envio_email; 
	`)
	fmt.Printf("Se borrarion las KEYS\n")
	if err != nil {
		log.Fatal(err)
	}
}

func insertarDatos(db *sql.DB) {
	result, err := db.Exec(`insert into paciente values (1, 'Juan', 'Perez', 12345678, '1978-05-08', 721, 523456, 'Suipacha 123', '+1153213421', 'juanperez1@gmail.com')
                        , (2, 'Maria','Rodriguez', 23456789, '1980-06-09', 722, 234567, 'Av. Libertador 123', '+1153213422', 'mariarodriguez1@gmail.com')
                        , (3, 'Pedro', 'Gomez', 34567890, '1982-07-10', 723, 345678, 'Calle 123', '+1153213423', 'pedrogomez1@gmail.com')
                        , (4, 'Lucia', 'Fernandez', 45678901, '1984-08-11', 723, 456789, 'Calle 456', '+1153213424', 'luciafernandez1@gmail.com')
                        , (5, 'Jorge', 'Gonzalez', 46789012, '1986-09-12', 722, 567890, 'Calle 789', '+1153213425', 'jorgegonzalez1@gmail.com')
                        , (6, 'Ana', 'Martinez', 37890123, '1988-10-13', 722, 678901, 'Calle 012', '+1153213426', 'anamartinez1@gmail.com')
                        , (7, 'Carlos', 'Sanchez', 28901234, '1990-11-14', 723, 789012, 'Calle 345', '+1153213427', 'carlossanchez1@gmail.com')
                        , (8, 'Laura', 'Romero', 19012345, '1992-12-15', 721, 890123, 'Calle 678', '+1153213428', 'lauraromero1@gmail.com')
                        , (9, 'Federico', 'Diaz', 20123456, '1994-01-16', 721, 901234, 'Calle 901', '+1153213429', 'federicodiaz1@gmail.com')
                        , (10, 'Mariana', 'Castro', 32345670, '1996-02-17', 722, 123459, 'Calle 234', '+1153213430', 'marianacastro1@gmail.com')
                        , (11, 'Roberto', 'Alvarez', 42345678, '1998-03-18', 721, 234569, 'Av. Libertador 456', '+1153213431', 'robertoalvarez1@gmail.com')
                        , (12, 'Sofia', 'Acosta', 33456789, '2000-04-19', 722, 345677, 'Calle 789', '+1153213432', 'sofiaacosta1@gmail.com')
                        , (13, 'Martin', 'Torres', 24567890, '2002-05-20', 723, 456784, 'Calle 012', '+1153213433', 'martintorres1@gmail.com')
                        , (14, 'Valentina', 'Ruiz', 15678901, '2004-06-21', 721, 567895, 'Calle 345', '+1153213434', 'valentinaruiz1@gmail.com')
                        , (15, 'Agustin', 'Sosa', 26789012, '2006-07-22', 722, 678905, 'Calle 678', '+1153213435', 'agustinsosa1@gmail.com')
                        , (16, 'Camila', 'Castro', 37890123, '2008-08-23', 723, 789015, 'Calle 901', '+1153213436', 'camilacastro1@gmail.com')
                        , (17, 'Lucas', 'Fernandez', 48901234, '2010-09-24', 721, 890125, 'Calle 234', '+1153213437', 'lucasfernandez1@gmail.com')
                        , (18, 'Sofia', 'Gonzalez', 39012345, '2012-10-25', 722, 901235, 'Calle 567', '+1153213438', 'sofiagonzalez1@gmail.com')
                        , (19, 'Luisa', 'Perez', 22345670, '2002-05-20', 723, 345666, 'Calle 345', '+1153213439', 'luisaperez1@gmail.com')
                        , (20, 'Miguel', 'Rodriguez', 13456780, '2004-06-21', 721, 678747, 'Calle 678', '+1153213440', 'miguelrodriguez1@gmail.com');`)
	fmt.Printf("Se insertaron los pacientes\n")
	if err != nil {
		log.Fatal(err)
	}
	rowsAffected, err := result.RowsAffected()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Se insertaron %d filas.\n", rowsAffected)
	result, err = db.Exec(`insert into medique values (44458951, 'Lara', 'Dolores', 'Traumatologo', 3000.50, '+1153223425')
                        , (44458952, 'Juan', 'Cardozo', 'Cardiologo', 5000.00, '+1153233426')
                        , (44458953, 'Sofia', 'Rodriguez', 'Pediatra', 2500.00, '+1153243427')
                        , (44458954, 'Martin', 'Gonzalez', 'Dermatologo', 4000.00, '+1153253428')
                        , (44458955, 'Ana', 'Fernandez', 'Ginecologo', 3500.00, '+1153216429')
                        , (44458956, 'Diego', 'Maradona', 'Oftalmologo', 3200.00, '+1153273430')
                        , (44458957, 'Valentina', 'Sanchez', 'Psiquiatra', 4500.00, '+1153813431')
                        , (44458958, 'Lucas', 'Garcia', 'Neurologo', 3800.00, '+1153213492')
                        , (44458959, 'Camila', 'Lopez', 'Oncologo', 4200.00, '+1153213733')
                        , (44458960, 'Mateo', 'Diaz', 'Endocrinologo', 3600.00, '+1153213634')
                        , (44458961, 'Agustina', 'Torres', 'Infectologo', 3300.00, '+1153289435')
                        , (44458962, 'Cristina', 'Kirchner', 'Reumatologo', 4100.00, '+1153218836')
                        , (44458963, 'Micaela', 'Castro', 'Nutricionista', 2800.00, '+1153671437')
                        , (44458964, 'Tomas', 'Romero', 'Urologo', 3400.00, '+1153213434')
                        , (44458965, 'Julieta', 'Gomez', 'Oncologo', 4400.00, '+1153211239')
                        , (44458966, 'Luciana', 'Pereyra', 'Cardiologo', 4800.00, '+1153234540')
                        , (44458967, 'Ignacio', 'Gutierrez', 'Traumatologo', 3100.00, '+1123163441')
                        , (44458968, 'Florencia', 'Alvarez', 'Pediatra', 2900.00, '+1153211232')
                        , (44458969, 'Santiago', 'Rojas', 'Dermatologo', 3700.00, '+1153243243')
                        , (44458970, 'Lucia', 'Garcia', 'Oftalmologo', 3900.00, '+11532132344');`)
	fmt.Printf("Se insertaron los mediques\n")
	if err != nil {
		log.Fatal(err)
	}
	rowsAffected, err = result.RowsAffected()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Se insertaron %d filas.\n", rowsAffected)
	result, err = db.Exec(`insert into consultorio values (1, 'Consultorio 1', 'Calle Consultorio 1', '1619', '1138624484'),
														(2, 'Consultorio 2', 'Calle Consultorio 2', '1620', '1157489652'),
														(3, 'Consultorio 3', 'Calle Consultorio 3', '1621', '1198654712'),
														(4, 'Consultorio 4', 'Calle Consultorio 4', '1622', '1185479623'),
														(5, 'Consultorio 5', 'Calle Consultorio 5', '1623', '1135987462'),
														(6, 'Consultorio 6', 'Calle Consultorio 6', '1624', '1155847999'),
														(7, 'Consultorio 7', 'Calle Consultorio 7', '1625', '1152366487');
	`)
	fmt.Printf("Se insertaron consultorios\n")
	if err != nil {
		log.Fatal(err)
	}
	rowsAffected, err = result.RowsAffected()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Se insertaron %d filas.\n", rowsAffected)
	result, err = db.Exec(`insert into obra_social values (721, 'OSDE', 'Roberto', 'De Magallanes', 1154785299, 'contacto@osde.com.ar'),
														(722, 'Swiss Medical', 'Fernanda', 'Rodriguez', 1122447586, 'contacto@swissmedical.com.ar'),
														(723, 'Galeno', 'Federico', 'Campos', 1125334578, 'contacto@galenoargentina.com.ar');
	`)
	fmt.Printf("Se insertaron las obras sociales\n")
	if err != nil {
		log.Fatal(err)
	}
	rowsAffected, err = result.RowsAffected()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Se insertaron %d filas.\n", rowsAffected)
	
	result, err = db.Exec(`insert into cobertura values (44458951, 721, 6000.00, 9000.00),
														(44458952, 722, 20000.00, 80000.00),
														(44458953, 723, 1000.00, 4000.00),
														(44458954, 723, 4500.00, 45500.00),
														(44458955, 722, 1500.00, 13500.00),
														(44458956, 722, 25000.00, 45000.00),
														(44458957, 723, 3000.00, 0.00),
														(44458958, 721, 12000.00, 78000.00),
														(44458959, 721, 20000.00, 80000.00),
														(44458960, 722, 4000.00, 56000.00),
														(44458961, 721, 0.00, 3300.00),
														(44458962, 722, 0.00, 12000.00),
														(44458963, 723, 5000.00, 0.00),
														(44458964, 721, 2500.00, 7500.00),
														(44458965, 722, 7500.00, 925000.00),
														(44458966, 723, 15000.00, 85000.00),
														(44458967, 722, 5000.00, 10000.00),
														(44458968, 722, 0.00, 5000.00),
														(44458969, 723, 4500.00, 5500.00),
														(44458970, 721, 70000.00, 0.00);
	`)
	fmt.Printf("Se insertaron las coberturas\n")
	if err != nil {
		log.Fatal(err)
	}
	rowsAffected, err = result.RowsAffected()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Se insertaron %d filas.\n", rowsAffected)
	
	result, err = db.Exec(`insert into agenda values (44458951, 1, 1, '06:00', '12:00', '1 hour')
					, (44458951, 3, 1, '06:00', '12:00', '1 hour')
					, (44458952, 2, 3, '06:00', '12:00', '30 minutes')
					, (44458952, 5, 3, '06:00', '12:00', '30 minutes')
					, (44458953, 3, 5, '06:00', '12:00', '15 minutes')
					, (44458953, 6, 5, '06:00', '12:00', '15 minutes')
					, (44458954, 7, 4, '06:00', '12:00', '30 minutes')
					, (44458955, 5, 1, '12:00', '18:00', '1 hour')
					, (44458956, 6, 6, '12:00', '18:00', '1 hour')
					, (44458957, 3, 2, '12:00', '18:00', '30 minutes')
					, (44458958, 4, 3, '12:00', '18:00', '15 minutes')
					, (44458959, 1, 4, '12:00', '18:00', '40 minutes')
					, (44458960, 2, 5, '12:00', '18:00', '1 hour')
					, (44458961, 4, 1, '18:00', '00:00', '1 hour')
					, (44458962, 5, 2, '18:00', '00:00', '30 minutes')
					, (44458963, 2, 3, '18:00', '00:00', '15 minutes')
					, (44458964, 3, 4, '18:00', '00:00', '30 minutes')
					, (44458965, 1, 5, '18:00', '00:00', '40 minutes')
					, (44458966, 2, 1, '00:00', '06:00', '30 minutes')
					, (44458967, 3, 2, '00:00', '06:00', '1 hour')
					, (44458968, 4, 3, '00:00', '06:00', '1 hour')
					, (44458969, 5, 4, '00:00', '06:00', '30 minutes')
					, (44458970, 1, 5, '00:00', '06:00', '1 hour');
	`)
	fmt.Printf("Se insertaron los datos en la agenda\n")
	if err != nil {
		log.Fatal(err)
		}
	rowsAffected, err = result.RowsAffected()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Se insertaron %d filas.\n", rowsAffected)
    
    result, err = db.Exec(`insert into solicitud_reservas values (1, 1, 44458951, '2023-11-30', '08:00:00'),
																(2, 2, 44458952, '2023-11-15', '11:00:00'),
																(3, 3, 44458953, '2023-11-01', '07:15:00'),
																(4, 4, 44458954, '2023-12-15', '06:30:00'),
																(5, 5, 44458955, '2023-12-31', '12:00:00'),
																(6, 6, 44458956, '2024-01-02', '16:00:00'),
																(7, 7, 44458957, '2024-02-13', '15:30:00'),
																(8, 8, 44458958, '2024-04-15', '17:45:00'),
																(9, 9, 44458959, '2024-03-11', '13:40:00'),
																(10, 10, 44458960, '2024-05-07', '14:00:00'),
																(11, 11, 44458961, '2024-01-19', '18:00:00'),
																(12, 12, 44458962, '2023-11-28', '19:30:00'),
																(13, 13, 44458963, '2023-12-17', '23:45:00'),
																(14, 14, 44458964, '2023-12-16', '20:00:00'),
																(15, 15, 44458965, '2024-01-30', '21:20:00'),
																(16, 16, 44458966, '2024-01-24', '02:30:00'),
																(17, 17, 44458967, '2023-12-25', '04:00:00'),
																(18, 18, 44458968, '2024-02-28', '05:30:00'),
																(19, 19, 44458969, '2024-04-01', '04:30:00'),
																(20, 20, 44458970, '2023-11-26', '00:00:00');
	`)
	fmt.Printf("Se insertaron las solicitudes de reserva\n")
	if err != nil {
		log.Fatal(err)
	}
	rowsAffected, err = result.RowsAffected()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Se insertaron %d filas.\n", rowsAffected)
}

func generarTurnos(db *sql.DB) {
	_, err := db.Exec(`create or replace function generar_turnos_disponibles(anio int, mes int)
					returns boolean as $$
					declare
						_dia int;
						_fecha timestamp;
						_nroturno int = 0;
						_dnis record;
						_mediquedni cursor for select dni_medique from agenda;
						_dias record;
						_agendadias cursor for select dia from agenda;
						_basura record;
					begin
						select t.fecha into _basura from turno t
						where extract(year from t.fecha) = anio and extract(month from t.fecha) = mes;
						if found then
							return false;
						else
							for _dnis in _mediquedni
							loop
								for _dias in _agendadias
								loop
							_nroturno := _nroturno + 1;
							_fecha = to_timestamp(anio || '-' || mes || '-' || _dias.dia, 'YYYY-MM-DD');
									insert into turno(nro_turno, fecha, dni_medique, estado) values (_nroturno, _fecha, _dnis.dni_medique, 'Disponible');
								end loop;
							end loop;
							return true;
						end if;
					end
					$$ language plpgsql;
	`)
	
	if err != nil {
		log.Fatal(err)
	}
    fmt.Printf("CREATE FUNCTION\n")
}
func crearTurnos(db *sql.DB) {
	var anio int
	var mes int	
	fmt.Printf("Ingresa año y mes separados con una , :  ")
	fmt.Scanf("%d, %d", &anio, &mes)
	rows, err := db.Query("select generar_turnos_disponibles($1, $2)", anio, mes)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next(){
		var resu bool
		err = rows.Scan(&resu)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%b", resu)
	}
}
func reservarFuncion(db *sql.DB) {
	_, err := db.Exec(`create or replace function reservar_turnos(_nro_paciente int, _dni_medique int, _fecha timestamp)
returns boolean as $$
declare
	_basura record;
begin
  select * from medique m into _basura where m.dni_medique in (_dni_medique);
  if found then
    select * from paciente p into _basura where p.nro_paciente in (_nro_paciente);
    if found then
      if exists (select * from paciente p, medique m, cobertura c where p.nro_obra_social = c.nro_obra_social and c.dni_medique = m.dni_medique and p.nro_paciente = _nro_paciente) then
        if exists (select * from turno t where t.fecha in (_fecha) and t.estado = 'Disponible') then
            select s.nro_paciente into _basura from solicitud_reservas s
            group by s.nro_paciente having count (s.nro_paciente) > 5;
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
          raise notice 'turno inexistente ó no disponible.';
          return false;
        end if;
      else
        raise notice 'obra social de paciente no atendida por le médique.';
        return false;
      end if;
    else
      raise notice 'nro de historia clínica no válido.';
      return false;
    end if;
  else
    raise notice 'dni de médique no válido.';
    return false;
  end if;
end
$$ language plpgsql;
	`)
	if err != nil {
		log.Fatal(err)
	}
    fmt.Printf("CREATE FUNCTION\n")
}
func reservarTurnos(db *sql.DB) {
	var _nro_paciente int
	var _dni_medique int
	var _fecha time.Time
	fmt.Printf("Ingresa numero de paciente, dni de medico y la fecha separados con una , :  ")
	fmt.Scanf("%d, %d, %v", &_nro_paciente, &_dni_medique, &_fecha)
	rows, err := db.Query("select reservar_turnos($1, $2, $3)", _nro_paciente, _dni_medique, _fecha)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next(){
		var resu bool
		err = rows.Scan(&resu)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%b", resu)
	}
}

/*func anularFuncion(db *sql.DB) {
	_, err := db.Exec(`create or replace function anular_turno(dni_mediques int, f_desde date, f_hasta date, nre_turno int) 
returns int as $$
declare
Cont_cancel int :=o;

begin

for turno in select * from turno where dni_mediques=dni_medique and f_reserva between f_desde and f_hasta;

loop

insert into reprogramacion
(id_paciente, f_cance, razon)
values(turno.id_paciente, current_date, 'cancelado razones del medico')

update into turno
set estado = 'Cancelado'
where nro_turno = nre_turno;

Cont_cancel:= Cont_cancel + 1;

end loop;

end;
$$language plpgsql;
	`)
	if err != nil {
		log.Fatal(err)
	}
fmt.Printf("CREATE FUNCTION\n")
}
func anularTurnos(db *sql.DB) {
	var dni_mediques int
	var f_desde string
	var f_hasta string
	var nre_turno int
	fmt.Printf("Ingresa dni de medico, fecha desde, fecha hasta y numero de turno separados por una , :  ")
	fmt.Scanf("%d, %s, %s, %d", &dni_mediques, &f_desde, &f_hasta, &nre_turno)
	rows, err := db.Query("select anular_turno($1, $2, $3, $4)", dni_mediques, f_desde, f_hasta, nre_turno)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next(){
		var resu int
		err = rows.Scan(&resu)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%d", resu)
	}
}
func atencionFuncion(db *sql.DB) {
	_, err := db.Exec(`create or replace function atender_turno(nro_turno_hg int) 
returns boolean as $$
declare
    v_estado_turno char(10); -- o lo hago con records?
    v_fecha_turno date;

begin

    -- comprueba si el turno existe
   
 select estado, fecha into v_estado_turno, v_fecha_turno
    from turno
    where  nro_turno = nro_turno_hg;

    
        -- acá sino encuentra entonces carga que el turno no existe

if not found then
        insert into error (f_turno, nro_error, operacion, f_error, motivo)
        values (current_timestamp, nro_turno_hg, 'atención', current_timestamp, 'nro de turno no válido');
        returns false;
    end if;

    -- se revisa si el turno esta reservado, si no lo esta, carga el turno como "turno no reservado"
    if v_estado_turno = 'reservado' then
 
        insert into error (f_turno, nro_error, operacion, f_error, motivo)
        values (current_timestamp, nro_turno_hg, 'atención', current_timestamp, 'turno no reservado');
        returns false;
    end if;

    -- comprueba si el turno es equivalente o igual al dia de la fecha sino lo es, entonces;  carga un error si el turno no equivale a la fecha del dia, este lo carga como " turno no correspondiente a la fecha del dia
    
    if v_fecha_turno == current_date then   -- en caso de no andar poner <> (estructura de comparación)
        insert into error (f_turno, nro_error, operacion, f_error, motivo)
        values (current_timestamo, nro_turno_hg, 'atención', current_timestamp, 'turno no corresponde a la fecha del día');
        returns false;
    end if;

    -- actualiza el estado a atendido
    update turno
    set estado = 'atendido'
    where nro_turno = nro_turno_hg;

    returns true;
end if;
$$ language plpgsql;
	`)
	if err != nil {
		log.Fatal(err)
	}
fmt.Printf("CREATE FUNCTION\n")
}
func atencionTurnos(db *sql.DB) {
	var nro_turno_hg int
	fmt.Printf("Ingresa numero de turno: ")
	fmt.Scanf("%d", &nro_turno_hg)
	rows, err := db.Query("select atender_turno($1)", nro_turno_hg)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next(){
		var resu bool
		err = rows.Scan(&resu)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%b", resu)
	}
}
func liqObraSocialFuncion(db *sql.DB) {
	_, err := db.Exec(`create or replace function generar_liquidacion_obra_social(
    in p_mes int,
    in p_anno int,
    in p_obra_social_id int
) returns table (
    paciente_historia_clinica_id int,
    monto_total numeric,
    detalle text
) as
$$
declare
    v_monto_total numeric := 0;
begin
  
    insert into liquidaciones_obra_social (mes, anno, obra_social_id)
    values (p_mes, p_anno, p_obra_social_id)
    returning id into v_liquidacion_id;


    return query
    update turnos
    set estado = 'liquidado'
    where fecha_hora >= to_date(p_mes || '-' || p_anno, 'MM-YYYY') and fecha_hora < to_date((p_mes + 1) || '-' || p_anno, 'MM-YYYY')
        and obra_social_id = p_obra_social_id
        and estado = 'reservado'
    returning
        paciente_historia_clinica_id,
        monto_consulta;

 
    select sum(monto_consulta) into v_monto_total
    from turnos
    where fecha_hora >= to_date(p_mes || '-' || p_anno, 'MM-YYYY') and fecha_hora < to_date((p_mes + 1) || '-' || p_anno, 'MM-YYYY')
        and obra_social_id = p_obra_social_id
        and estado = 'liquidado';


    update liquidaciones_obra_social
    set monto_total = v_monto_total
    where id = v_liquidacion_id;

    return next v_monto_total;
end;
$$ language plpgsql;`)
if err != nil {
	log.Fatal(err)
}
fmt.Printf("CREATE FUNCTION\n")

}
func generarLiqObraSocial(db *sql.DB) {
	defer db.Close()
	var p_mes int
    var p_anno int
    var p_obra_social_id int
	fmt.Printf("Ingresa mes, año y numero de obra social separados por una , : ")
	fmt.Scanf("%d, %d, %d", &p_mes, &p_anno, &p_obra_social_id)
	rows, err := db.Query("select generar_liquidacion_obra_social($1, $2, $3)", p_mes, p_anno, p_obra_social_id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next(){
		var resu bool
		err = rows.Scan(&resu)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%d", resu)
	}
}

func envioEmail(db *sql.DB) {
	_, err := db.Exec(`create or replace function enviar_email(_email text, _asunto text, _cuerpo text)
returns trigger as $$
declare
	_email text;
	_asunto text;
	_cuerpo text;
	_nro_email int = 0;
	_emails record;
	_emailpaciente cursor for select email_paciente from envio_email;
	_basura record;
begin
	for _emails in _emailpaciente
		select e.nro_email into _basura from envio_email e;
		if found then
			_nro_email:=_nro_email + nro_email;
			_nro_email:=_nro_email + 1;
			insert into envio_email 
			values (_nro_email, current_date, _email, _asunto, _cuerpo, current_date, 'Enviado');
			exit;
		else
			_nro_email := _nro_email + 1;
			insert into envio_email 
			values (_nro_email, current_date, _email, _asunto, _cuerpo, current_date, 'Enviado');
			exit;
		end if;
	end loop;
end
$$ language pspgsql;

create trigger email_reserva_turno
after update of estado on turno
for each row
when (estado = 'Reservado')
execute function enviar_email();

create trigger email_cancelar_turno();
after update of estado on turno
for each row
when (estado = 'Cancelado')
execute function enviar_email();
	`)
	if err != nil {
		log.Fatal(err)
	}
fmt.Printf("CREATE FUNCTION\n")
}
*/
