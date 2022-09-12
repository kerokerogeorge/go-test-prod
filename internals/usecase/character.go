package usecase

import (
	"strconv"

	"github.com/kerokerogeorge/go-gacha-api/internals/domain/model"
	"github.com/kerokerogeorge/go-gacha-api/internals/domain/repository"
)

type CharacterUsecase interface {
	List() ([]*model.Character, error)
	Get(characterId string) (*model.Character, error)
	Create(character *model.Character) (*model.Character, error)
	Delete(character *model.Character) error
	GetCharactersWithEmmitionRate(gachaId string) ([]*model.CharacterWithEmmitionRate, error)
	GetGachaCharacters(characterId string) ([]*model.CharacterEmmitionRate, error)
	DeleteGachaCharacters(gachaCharacters []*model.CharacterEmmitionRate) error
	GetUserCharacters(characterId string) ([]*model.UserCharacter, error)
	DeleteUserCharacters(userCharacters []*model.UserCharacter) error
}

type characterUsecase struct {
	characterRepo             repository.CharacterRepository
	characterEmmitionRateRepo repository.CharacterEmmitionRateRepository
	userCharcacterRepo        repository.UserCharcacterRepository
}

func NewCharacterUsecase(cr repository.CharacterRepository, cerr repository.CharacterEmmitionRateRepository, rr repository.UserCharcacterRepository) CharacterUsecase {
	return &characterUsecase{
		characterRepo:             cr,
		characterEmmitionRateRepo: cerr,
		userCharcacterRepo:        rr,
	}
}

func (cu *characterUsecase) List() ([]*model.Character, error) {
	return cu.characterRepo.GetCharacters()
}

func (cu *characterUsecase) Get(characterId string) (*model.Character, error) {
	id, _ := strconv.Atoi(characterId)
	return cu.characterRepo.GetCharacter(id)
}

func (cu *characterUsecase) Create(character *model.Character) (*model.Character, error) {
	return cu.characterRepo.CreateCharacter(character)
}

func (cu *characterUsecase) Delete(character *model.Character) error {
	return cu.characterRepo.DeleteCharacter(character)
}

func (cu *characterUsecase) GetCharactersWithEmmitionRate(gachaId string) ([]*model.CharacterWithEmmitionRate, error) {
	return cu.characterEmmitionRateRepo.GetCharacterWithEmmitionRate(gachaId)
}

func (cu *characterUsecase) GetGachaCharacters(characterId string) ([]*model.CharacterEmmitionRate, error) {
	return cu.characterEmmitionRateRepo.GetGachaCharactersFromCharacterId(characterId)
}

func (cu *characterUsecase) DeleteGachaCharacters(gachaCharacters []*model.CharacterEmmitionRate) error {
	for _, gachaCharacter := range gachaCharacters {
		err := cu.characterEmmitionRateRepo.DeleteGachaCharacter(gachaCharacter)
		if err != nil {
			return err
		}
	}

	return nil
}

func (cu *characterUsecase) GetUserCharacters(characterId string) ([]*model.UserCharacter, error) {
	return cu.userCharcacterRepo.GetUserCharacters(characterId, "CHARACTER")
}

func (cu *characterUsecase) DeleteUserCharacters(userCharacters []*model.UserCharacter) error {
	for _, userCharacter := range userCharacters {
		err := cu.userCharcacterRepo.DeleteUserCharacter(userCharacter)
		if err != nil {
			return err
		}
	}

	return nil
}
