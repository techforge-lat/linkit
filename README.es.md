# Lenguajes

- [English](./README.md)
- [Español](./README.es.md)

# Linkit

**Linkit** es un contenedor de inyección de dependencias minimalista para Go. Simplifica la gestión de tipos y dependencias en proyectos de Go al recolectar e inicializar tipos.

## Motivación

En proyectos anteriores, la gestión de dependencias carecía de un enfoque integral y estandarizado, lo que generaba confusión, especialmente para los recién llegados. Cada proyecto seguía patrones diferentes, lo que llevaba a inconsistencias y complejidades en la gestión de dependencias entre equipos.

**Linkit** resuelve esto al proporcionar una manera clara y definida de manejar dependencias, asegurando un enfoque consistente y estructurado que es fácil de entender y mantener. Reduce la confusión, acelera la incorporación de nuevos miembros y simplifica el proceso de inyección de dependencias para todos los desarrolladores.

## Conceptos

- **Contenedor**: Una colección de tipos inicializados.
- **Dependencia Principal**: La dependencia principal en un contexto (manejador, caso de uso, repositorio, etc.).
- **Dependencia Auxiliar**: Dependencias principales que son requeridas por otras dependencias principales.

### Cómo usarlo

Linkit proporciona un flujo de trabajo de 3 pasos.

1. **Inicializar el contenedor**: Crear una instancia del contenedor para gestionar las dependencias de tu aplicación.
```go
    container := linkit.New()
```

2. **Proveer dependencias**: Registrar dependencias utilizando el método Provide. Esto registra una función que proporcionará la dependencia cuando sea necesaria.
```go
	userUseCase := user.NewUseCase(user.NewPsqlRepository())
	userHandler := user.NewHandler(userUseCase)

    // Para proporcionar, necesitamos definir una convención de nomenclatura para cada nombre de dependencia
    // aquí estamos usando `nombre-del-módulo.capa`= 'user.usecase'
    // estos nombres se pueden definir en un solo archivo
	container.Provide(linkit.DependencyName("user.usecase"), userUseCase)
	container.Provide(linkit.DependencyName("user.handler"), userHandler)

	// debe ser después de que se hayan proporcionado todas las demás dependencias principales
	// esto ejecutará cada método ResolveAuxiliaryDependencies de cada dependencia
	if err := container.ResolveAuxiliaryDependencies(); err != nil {
		return nil, err
	}
```

3. **Resolver dependencias**: Utilizar la función `Resolve` para recuperar una dependencia proporcionada.
```go
    // UserUseCase es una dependencia principal
    type UserUseCase struct {
        // repository es una dependencia auxiliar que se establece por el constructor
        repository Repository

        // role es una dependencia auxiliar que se establece por el método ResolveAuxiliaryDependencies requerido por linkit
        role       RoleUseCase
    }

    // NewUseCase inicializa la dependencia principal
    func NewUseCase(repository Repository) *UserUseCase {
        return &UserUseCase{
            repository: repository,
        }
    }

    // ResolveAuxiliaryDependencies establece las dependencias auxiliares,
    // este método es llamado por el método ResolveAuxiliaryDependencies del contenedor linkit creado en el primer paso
    // después de que se ha proporcionado cada dependencia principal para que no encontremos un error al resolverlas aquí
    func (u UserUseCase) ResolveAuxiliaryDependencies(container *linkit.DependencyContainer) error {
        roleUseCase, err := linkit.Resolve[RoleUseCase](container, linkit.DependencyName("role.usecase"))
        if err != nil {
            return err
        }
        u.role = roleUseCase

        return nil
    }
```

*consulta la carpeta [examples/](./examples/) para más detalles.*

## Limitaciones o Problemas Conocidos

- **Requerimiento de Punteros**: Las dependencias deben pasarse como punteros, lo que puede requerir un cuidado adicional al definirlas y usarlas.
- **Nomenclatura Manual**: Cada dependencia debe proporcionarse con un nombre. Si bien esto permite flexibilidad, puede dar lugar a errores de nomenclatura, ya que los nombres no están "tipados" o impuestos por el compilador. Si intentas resolver una dependencia con el nombre incorrecto, se generará un error.

