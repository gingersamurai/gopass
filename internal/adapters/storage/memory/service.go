package memory

import (
	"fmt"
	"gopass/internal/adapters/storage"
	"gopass/internal/entity"
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
		return entity.Service{}, fmt.Errorf("memoryStorage.GetService(): %w", storage.ErrServiceNotFound)
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

	return entity.Service{}, fmt.Errorf("memoryStorage.GetServiceByName(): %w", storage.ErrServiceNotFound)
}

func (ms *MemoryStorage) DeleteService(id int64) error {
	ms.Lock()
	defer ms.Unlock()

	if _, ok := ms.serviceData[id]; !ok {
		return fmt.Errorf("memoryStorage.GetService(): %w", storage.ErrServiceNotFound)
	}
	delete(ms.serviceData, id)
	return nil
}
