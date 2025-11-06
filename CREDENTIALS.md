# 🔐 Credenciales de Desarrollo

## Usuarios de Prueba

### 👨‍💼 Administrador
- **Usuario:** `admin`
- **Contraseña:** `admin`
- **Email:** `admin@gmail.com`
- **Rol:** Administrador (acceso completo)
- **Permisos:** Todos los permisos del sistema

### 👨‍⚕️ Médico 1
- **Usuario:** `pcastejon`
- **Contraseña:** `1234`
- **Email:** `pcastejon@gmail.com`
- **Rol:** Médico
- **Permisos:** Ver/gestionar pacientes, consultas, exámenes, citas

### 👨‍⚕️ Médico 2
- **Usuario:** `jmerida`
- **Contraseña:** `1234`
- **Email:** `jmerida@gmail.com`
- **Rol:** Administrador
- **Permisos:** Acceso completo

### 👩‍💼 Secretaria
- **Usuario:** `jlopez`
- **Contraseña:** `1234`
- **Email:** `jlopez@gmail.com`
- **Rol:** Secretaria
- **Permisos:** Ver/gestionar pacientes y citas

## 🌐 URLs de Acceso

- **Frontend:** http://localhost:5173
- **Backend API:** http://localhost:8080
- **Login:** http://localhost:5173/login

## 📋 Instrucciones de Uso

1. Accede a http://localhost:5173/login
2. Ingresa las credenciales de cualquier usuario de arriba
3. Haz clic en "Ingresar"
4. Serás redirigido al dashboard automáticamente

## 🏥 Datos de Prueba

La base de datos incluye:
- **19 pacientes** (10 adultos, 9 menores con tutores)
- **Consultas** con diagnósticos y tratamientos
- **Exámenes** médicos (algunos pendientes, otros completados)
- **Citas** programadas
- **Horarios** laborales y especiales

### Pacientes de Ejemplo
- **ID 1:** Juan Pérez Martínez
- **ID 2:** María González López
- **ID 3:** Carlos Sánchez Díaz
- **ID 4:** Ana Fernández Ruiz
- (y 15 más...)

## ⚠️ Notas Importantes

- Estas son credenciales de **DESARROLLO** únicamente
- **NUNCA** uses estas credenciales en producción
- La base de datos se resetea cada vez que se ejecuta `docker-compose down -v`
- El token JWT expira después de 100,000 horas (~11.4 años) en desarrollo

## 🔧 Troubleshooting

### Error 401 (Unauthorized)
- **Causa**: No estás autenticado
- **Solución**: Ve a http://localhost:5173/login e inicia sesión

### Error 403 (Acceso denegado)
- **Causa**: Tu rol no tiene los permisos necesarios
- **Solución**: Usa una cuenta con permisos suficientes (ej: `admin`)
- **Nota**: El panel de administración requiere los permisos `manejar-roles` y `manejar-usuarios`

### Página en blanco
- **Causa**: Los contenedores no están corriendo
- **Solución**: Ejecuta `docker-compose -f docker-compose.dev.yml up -d`
