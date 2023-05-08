package memory_storage

import (
	"fmt"
	"gopass/internal/entity"
	"gopass/internal/usecase"
)

func (ms *MemoryStorage) AddService(service entity.Service) (int64, error) {
	ms.Lock()
	defer ms.Unlock()

	service.Id = ms.nextServiceId
	ms.nextServiceId++

	ms.serviceData[service.Id] = service
	return service.Id, nil
}

func (ms *MemoryStorage) GetService(id int64) (entity.Service, error) {
	ms.RLock()
	defer ms.RUnlock()

	if _, ok := ms.serviceData[id]; !ok {
		return entity.Service{}, fmt.Errorf("memoryStorage.GetService(): %w", usecase.ErrServiceNotFound)
	}

	return ms.serviceData[id], nil
}

func (ms *MemoryStorage) GetServiceByName(name string) (entity.Service, error) {
	ms.RLock()
	defer ms.RUnlock()

	for _, service := range ms.serviceData {
		if service.Name == name {
			return service, nil
		}
	}

	return entity.Service{}, usecase.ErrServiceNotFound
}

func (ms *MemoryStorage) DeleteService(id int64) error {
	ms.Lock()
	defer ms.Unlock()

	if _, ok := ms.serviceData[id]; !ok {
		return fmt.Errorf("memoryStorage.GetService(): %w", usecase.ErrServiceNotFound)
	}
	delete(ms.serviceData, id)
	return nil
}
