package nightmare

type Specification interface {
	isSatisfiedBy(nm *Nightmare) bool
}
type ExistSpecification struct {
	Repo Repository
}

func (es *ExistSpecification) Init( r Repository)  {
	es.Repo = r
}
func (es *ExistSpecification) isSatisfiedBy(nightmare *Nightmare) bool  {

	nm , err := es.Repo.GetNightmare(nightmare)

	if err != nil {
		return false
	}
	if nm.NameEN == ""{
		return false
	}

	return true
}
