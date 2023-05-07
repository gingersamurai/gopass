package memory_storage

import "gopass/internal/entity"

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
		return entity.Service{}, errServiceNotFound
	}

	return ms.serviceData[id], nil
}

func (ms *MemoryStorage) DeleteService(id int64) error {
	ms.Lock()
	defer ms.Unlock()

	if _, ok := ms.serviceData[id]; !ok {
		return errServiceNotFound
	}
	delete(ms.serviceData, id)
	return nil
}
